// Copyright (c) 2021 Dave Marsh. See LICENSE.

package camera

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
	"time"

	"github.com/centretown/tiny-fabb/forms"
	"github.com/centretown/tiny-fabb/servo"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/mattn/go-mjpeg"
)

type Camera struct {
	ID           string        `json:"id"`
	Url          string        `json:"url"`
	Title        string        `json:"title"`
	Active       bool          `json:"active"`
	Interval     time.Duration `json:"interval"`
	ServoIndeces []uint        `json:"servoIndeces"`
	Streamer     Streamer      `json:"streamer"`

	Forms  forms.Forms    `json:"-"`
	Servos []*servo.Servo `json:"-"`

	layout *template.Template
	wg     *sync.WaitGroup
	stream *mjpeg.Stream
}

type Cameras map[string]*Camera

func (cam *Camera) Setup(router *mux.Router, servos *servo.Connector) {
	cam.Forms = cam.Streamer.BindProperties()
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
	go cam.poll()
}

func (cam *Camera) Stop() {
	cam.stream.Close()
	cam.wg.Wait()
}

func (cam *Camera) poll() {
	defer cam.wg.Done()
	var (
		err               error
		repeat, threshold uint
	)
	threshold = 10

	var buf []byte
	for !cam.stream.Closed() {
		buf, err = cam.Streamer.Read()
		if err != nil {
			repeat++
			glog.Infoln(cam.ID, err, repeat)
			if repeat > threshold {
				break
			}
			continue
		}
		err = cam.stream.Update(buf)
		if err != nil {
			glog.Infoln(cam.ID, err)
			continue
		}
	}
}

func (cam *Camera) Apply(w http.ResponseWriter, r *http.Request) (err error) {
	id := forms.GetRequestString(r, "id")
	val := forms.GetRequestString(r, "val")
	err = cam.Set(id, val)
	if err != nil {
		forms.WriteError(w, err)
		return
	}
	return
}

func (cam *Camera) Get(id string) (val string, err error) {
	webid := forms.ToWebId(id)
	frm, ok := cam.Forms[webid]
	if !ok {
		err = fmt.Errorf("id '%s' not found", id)
		return
	}
	val = fmt.Sprint(frm.Value)
	return
}

func (cam *Camera) Set(id string, val string) (err error) {
	webid := forms.ToWebId(id)
	frm, ok := cam.Forms[webid]
	if !ok {
		err = fmt.Errorf("id '%s' not found", id)
		return
	}

	ent := frm.Entries[0]
	if !ent.InBounds(val) {
		err = fmt.Errorf("%s: %s out of range", id, val)
		return
	}

	err = ent.ScanInput(val, frm.Value)
	if err != nil {
		return
	}

	err = cam.Streamer.SetProperty(ent, val)
	return
}
