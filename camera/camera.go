package camera

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/mattn/go-mjpeg"
	"gocv.io/x/gocv"
)

type Camera struct {
	Title      string
	Settings   CameraStatus
	wg         *sync.WaitGroup
	stream     *mjpeg.Stream
	streamUrl  string
	controlUrl string
	statusUrl  string
	captureUrl string
	interval   time.Duration
}

func (cam *Camera) ShowWindow() func(img image.Image) {
	resize := func(window *gocv.Window, img image.Image) {
		b := img.Bounds()
		window.ResizeWindow(b.Max.X-b.Min.X, b.Max.Y-b.Min.Y)
	}
	resized := false
	window := gocv.NewWindow(cam.streamUrl)

	return func(img image.Image) {
		if !resized {
			resize(window, img)
		}
		// if !window.IsOpen() {
		// }

		mat, err := gocv.ImageToMatRGB(img)
		if err != nil {
			fmt.Println(err)
			return
		}
		if mat.Empty() {
			fmt.Println("matrix empty")
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

	dec, err := mjpeg.NewDecoderFromURL(cam.streamUrl)
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
				break
			}
			glog.Info(cam.Title, err)
			continue
		}

		if viewer != nil {
			viewer(img)
		}

		buf.Reset()
		err = jpeg.Encode(&buf, img, nil)
		if err != nil {
			fmt.Println(cam.Title, err)
			continue
		}
		err = cam.stream.Update(buf.Bytes())
		if err != nil {
			fmt.Println(cam.Title, err)
			continue
		}
	}
}

func (cam *Camera) Start(router *mux.Router) {
	cam.bindSettings()
	cam.wg = &sync.WaitGroup{}
	cam.wg.Add(1)
	cam.stream = mjpeg.NewStreamWithInterval(cam.interval)

	go cam.proxy(cam.ShowWindow())
	prefix := "/" + cam.Title + "/"

	router.HandleFunc(prefix+"jpeg",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/jpeg")
			w.Write(cam.stream.Current())
		})

	router.HandleFunc(prefix+"mjpeg", cam.stream.ServeHTTP)
}

func (cam *Camera) Stop() {
	cam.stream.Close()
	cam.wg.Wait()
}
