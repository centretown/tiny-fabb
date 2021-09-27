package camera

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Cameras map[string]*Camera

func (cams Cameras) Start(router *mux.Router,
	layout *template.Template,
	interval time.Duration, urls ...string) {
	// LoadClassifier()

	router.HandleFunc("/camera-bar/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := layout.Lookup("camera-bar")
		if tmpl == nil {
			return
		}
		tmpl.Execute(w, cams)
	})

	for i, url := range urls {
		cam := &Camera{
			Name:       fmt.Sprintf("camera%d", i),
			Title:      fmt.Sprintf("Camera: %0d", i),
			ControlUrl: url + ":80/action",
			CaptureUrl: url + ":80/capture",
			StatusUrl:  url + ":80/status",
			StreamUrl:  url + ":81/stream",
			Interval:   interval,
			layout:     layout,
		}
		cams[cam.Name] = cam
		fmt.Println("SetFrameSize")
		cam.Set("frameSize", "8")
		fmt.Println("SetQuality")
		cam.Set("quality", "10")
		fmt.Println("GetStatus")
		cam.GetStatus()
		fmt.Println("Start")
		cam.Start(router)
	}
}

func (cams Cameras) Stop() {
	for _, cam := range cams {
		cam.Stop()
	}
	// CloseClassifier()
}
