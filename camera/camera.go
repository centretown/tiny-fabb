// Copyright (c) 2021 Dave Marsh. See LICENSE.

package camera

import (
	"bytes"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/centretown/tiny-fabb/forms"
	"github.com/centretown/tiny-fabb/servo"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/mattn/go-mjpeg"
	"gocv.io/x/gocv"
)

type Camera struct {
	ID           string         `json:"id"`
	Url          string         `json:"url"`
	Title        string         `json:"title"`
	Active       bool           `json:"active"`
	Settings     CameraStatus   `json:"settings"`
	ControlUrl   string         `json:"controlUrl"`
	StreamUrl    string         `json:"streamUrl"`
	StatusUrl    string         `json:"statusUrl"`
	CaptureUrl   string         `json:"captureUrl"`
	Interval     time.Duration  `json:"interval"`
	ServoIndeces []uint         `json:"servoIndeces"`
	Forms        forms.Forms    `json:"-"`
	Servos       []*servo.Servo `json:"-"`
	layout       *template.Template
	wg           *sync.WaitGroup
	stream       *mjpeg.Stream
}

type Cameras map[string]*Camera

func (cam *Camera) Setup(router *mux.Router, servos *servo.Connector) {
	cam.bindSettings()
	// go cam.proxy(cam.ShowWindow())
	prefix := "/" + cam.ID + "/"
	router.HandleFunc(prefix+"apply/{id}/{val}/",
		func(w http.ResponseWriter, r *http.Request) {
			cam.Apply(w, r)
		})

	router.HandleFunc(prefix+"pan-tilt/{pan}/{tilt}/{speed}/",
		func(w http.ResponseWriter, r *http.Request) {
			cam.PanTilt(w, r)
		})

	var err error
	cam.Servos, err = servos.Connect(cam.ServoIndeces...)
	if err != nil {
		glog.Warningln(err)
	}

	for _, svo := range cam.Servos {
		svo.Connect(router, prefix)
	}

	router.HandleFunc(prefix+"camera-full/",
		func(w http.ResponseWriter, r *http.Request) {
			tmpl := cam.layout.Lookup("camera-full")
			if tmpl == nil {
				return
			}
			tmpl.Execute(w, cam)
		})
	router.HandleFunc(prefix+"camera-settings/",
		func(w http.ResponseWriter, r *http.Request) {
			tmpl := cam.layout.Lookup("camera-settings")
			if tmpl == nil {
				return
			}
			tmpl.Execute(w, cam)
		})
	router.HandleFunc(prefix+"servo-settings/",
		func(w http.ResponseWriter, r *http.Request) {
			tmpl := cam.layout.Lookup("servo-settings")
			if tmpl == nil {
				return
			}
			tmpl.Execute(w, cam)
		})
	cam.Start()
	router.HandleFunc(prefix+"jpeg",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/jpeg")
			w.Write(cam.stream.Current())
		})
	router.HandleFunc(prefix+"mjpeg", cam.stream.ServeHTTP)
}

func (cam *Camera) Start() {
	cam.wg = &sync.WaitGroup{}
	cam.wg.Add(1)
	cam.stream = mjpeg.NewStreamWithInterval(cam.Interval)
	go cam.proxy(nil)
}

func (cam *Camera) Stop() {
	cam.stream.Close()
	cam.wg.Wait()
}

func (cam *Camera) PanTilt(w http.ResponseWriter, r *http.Request) {
	var (
		move = func(svo *servo.Servo, angle int, speed uint) {
			if angle >= 0 && angle <= 180 {
				svo.Command = servo.ServoMove
				svo.Angle = uint(angle)
				svo.Speed = speed
				svo.Apply(w)
			}
		}
	)
	if len(cam.Servos) < 1 {
		return
	}

	speed := forms.GetRequestUint(r, "speed")
	if speed < 1 || speed > 255 {
		speed = 50
	}

	change := forms.GetRequestInt(r, "pan")
	svo := cam.Servos[0]
	move(svo, change+int(svo.Angle), speed)

	if len(cam.Servos) < 2 {
		return
	}
	change = forms.GetRequestInt(r, "tilt")
	svo = cam.Servos[1]
	move(svo, change+int(svo.Angle), speed)
}

func (cam *Camera) ShowWindow() func(img image.Image) {
	resized := false
	resize := func(window *gocv.Window, img image.Image) {
		b := img.Bounds()
		window.ResizeWindow(b.Max.X-b.Min.X, b.Max.Y-b.Min.Y)
		resized = true
	}
	window := gocv.NewWindow(cam.StreamUrl)

	return func(img image.Image) {
		if !resized {
			resize(window, img)
		}
		// if !window.IsOpen() {
		// }

		mat, err := gocv.ImageToMatRGB(img)
		if err != nil {
			glog.Infoln(err)
			return
		}
		if mat.Empty() {
			glog.Infoln("matrix empty")
			return
		}
		// blue := color.RGBA{0, 0, 255, 0}
		// rects := classifier.DetectMultiScale(mat)
		// fmt.Printf("found %d faces\n", len(rects))
		// draw a rectangle around each face on the original image
		// for _, r := range rects {
		// 	gocv.Rectangle(&mat, r, blue, 3)
		// 	gocv.PutText(&mat, "ape",
		// 		image.Pt(r.Min.X, r.Min.Y),
		// 		gocv.FontItalic,
		// 		2.0,
		// 		color.RGBA{R: 255, A: 63}, 4)
		// }

		window.IMShow(mat)
		window.WaitKey(1)
	}
}

func (cam *Camera) proxy(viewer func(img image.Image)) {
	defer cam.wg.Done()

	dec, err := mjpeg.NewDecoderFromURL(cam.StreamUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	var buf bytes.Buffer
	for !cam.stream.Closed() {

		img, err := dec.Decode()
		if err != nil {
			msg := err.Error()
			if strings.Contains(msg, "connection reset by peer") ||
				strings.Contains(msg, "connection timed out") {
				glog.Error(err)
				return
			}
			glog.Info(cam.ID, err)
			continue
		}

		if viewer != nil {
			viewer(img)
		}

		buf.Reset()
		err = jpeg.Encode(&buf, img, nil)
		if err != nil {
			glog.Infoln(cam.ID, err)
			continue
		}
		err = cam.stream.Update(buf.Bytes())
		if err != nil {
			glog.Infoln(cam.ID, err)
			continue
		}
	}
}
