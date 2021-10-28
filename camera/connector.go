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
		cam     *Camera
		ok      bool
		profile *Profile
		err     error
	)

	for _, url := range urls {
		profile = conn.profile(url)
		cam, ok = conn.Cameras[profile.ID]
		if !ok {
			cam = &Camera{
				ID:           profile.ID,
				Url:          url,
				Title:        profile.Title,
				ServoIndeces: []uint{0, 1},
			}
			conn.Cameras[cam.ID] = cam
		}

		cam.layout = layout
		cam.Interval = conn.interval
		cam.Streamer = profile.streamer
		cam.Forms = cam.Streamer.BindProperties()

		err = cam.Streamer.Open()
		if err != nil {
			glog.Warningf("failed to open camera %s:%s -> %v\n", cam.ID, url, err)
			continue
		}

		err = cam.Streamer.UpdateProperties()
		if err != nil {
			glog.Warningf("failed to update camera %s:%s -> %v\n", cam.ID, url, err)
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

func (conn *Connector) profile(url string) (profile *Profile) {
	var i int
	profile = &Profile{
		Url: url,
	}
	if strings.HasPrefix(url, "http") {
		i = strings.LastIndex(url, ".")
		suffix := url[i+1:]
		profile.ID = fmt.Sprintf("camera%s", suffix)
		profile.Title = fmt.Sprintf("Camera: %s", suffix)
		st := &EspCam{}
		st.ControlUrl = url + ":80/action"
		st.CaptureUrl = url + ":80/capture"
		st.StatusUrl = url + ":80/status"
		st.StreamUrl = url + ":81/stream"
		profile.streamer = st
	} else {
		i := strings.LastIndex(url, "/")
		profile.ID = url[i+1:]
		profile.Title = strings.Title(profile.ID)
		st := &LocalCam{}
		st.StreamUrl = url
		profile.streamer = st
	}
	return
}
