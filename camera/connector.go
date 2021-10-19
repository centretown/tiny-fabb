package camera

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/centretown/tiny-fabb/servo"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

type Connector struct {
	Cameras    Cameras
	dataSource string
	layout     *template.Template
	interval   time.Duration
}

// LoadClassifier()
// CloseClassifier()

func NewConnector(dataSource string, layout *template.Template,
	interval time.Duration) (conn *Connector) {
	conn = &Connector{
		Cameras:    make(map[string]*Camera),
		dataSource: dataSource,
		layout:     layout,
		interval:   interval,
	}
	conn.Load(dataSource)
	return
}

func (conn *Connector) Start(router *mux.Router,
	servoController string,
	urls ...string) {

	layout := conn.layout
	servos := servo.NewConnector(servoController,
		conn.dataSource, layout, 9)

	router.HandleFunc("/camera-bar/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := layout.Lookup("camera-bar")
		if tmpl == nil {
			return
		}
		tmpl.Execute(w, conn.Cameras)
	})

	var (
		cam *Camera
		ok  bool
	)

	for _, url := range urls {
		i := strings.LastIndex(url, ".")
		if i == -1 {
			continue
		}
		suffix := url[i+1:]
		id := fmt.Sprintf("camera%s", suffix)
		cam, ok = conn.Cameras[id]
		if !ok {
			cam = &Camera{
				ID:           id,
				Url:          url,
				Title:        fmt.Sprintf("Camera: %s", suffix),
				ControlUrl:   url + ":80/action",
				CaptureUrl:   url + ":80/capture",
				StatusUrl:    url + ":80/status",
				StreamUrl:    url + ":81/stream",
				Interval:     conn.interval,
				ServoIndeces: []uint{0, 1},
			}
			conn.Cameras[cam.ID] = cam
		}
		cam.layout = layout
		err := cam.GetStatus()
		if err != nil {
			glog.Warningf("failed to connect camera %s:%s -> %v\n", cam.ID, url, err)
			continue
		}
		cam.Active = true
		cam.Setup(router, servos)
	}
	conn.Save(conn.dataSource)
}

func (conn *Connector) Load(dataSource string) (err error) {
	var b []byte
	b, err = ioutil.ReadFile(dataSource + "cameras.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &conn.Cameras)
	return
}

func (conn *Connector) Save(dataSource string) (err error) {
	var b []byte
	b, err = json.MarshalIndent(conn.Cameras, "", "  ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(dataSource+"cameras.json", b, 0640)
	return
}

func (conn *Connector) Stop() {
	for _, cam := range conn.Cameras {
		cam.Stop()
	}
}
