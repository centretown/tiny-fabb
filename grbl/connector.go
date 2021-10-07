package grbl

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"

	"github.com/centretown/tiny-fabb/camera"
	"github.com/centretown/tiny-fabb/monitor"
)

type Connector struct {
	Controllers map[string]*Controller
	dataSource  string
	layout      *template.Template
	cameras     camera.Cameras
}

func NewConnector(dataSource string, layout *template.Template, cameras camera.Cameras) (conn *Connector) {
	conn = &Connector{
		Controllers: make(map[string]*Controller),
		dataSource:  dataSource,
		layout:      layout,
		cameras:     cameras,
	}
	conn.Load()
	return
}

func (conn *Connector) Connect(bus *monitor.Bus) (ctl monitor.Controller, err error) {

	var (
		terminators     = []string{"ok", "error"}
		results         []string
		profile         *Profile
		gctl            *Controller
		ok              bool
		controllerCount int = len(conn.Controllers)
	)

	results, err = bus.Capture("$I", terminators...)
	if err != nil {
		return
	}

	profile, err = CheckProfile(results)
	if err != nil {
		return
	}

	gctl, ok = conn.Controllers[profile.ID]
	if ok {
		gctl.Profile = profile
		gctl.initialize(bus, conn.layout)
	} else {
		controllerCount++
		title := fmt.Sprintf("%s-%d", profile.Prefix(), controllerCount)
		gctl = NewController(bus, conn.layout)
		gctl.Profile = profile

		if len(profile.ID) == 0 {
			profile.ID = title
			s := fmt.Sprintf("$I=%s%s", ID, profile.ID)
			_, err = bus.Capture(s, terminators...)
			if err != nil {
				return
			}
		}
		gctl.ID = profile.ID
		if profile.IsESP32() {
			gctl.Title = title
		} else {
			gctl.Title = profile.ID
		}
		conn.Add(gctl)
	}
	// gctl.Cameras = conn.cameras

	for _, c := range gctl.CameraIDs {
		gctl.Cameras[c] = conn.cameras[c]
	}

	ctl = gctl
	err = gctl.startup()
	return
}

func (conn *Connector) Save() (err error) {
	var b []byte
	b, err = json.MarshalIndent(&conn.Controllers, "", "  ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(conn.dataSource+"grbl.json", b, 0640)
	return
}

func (conn *Connector) Load() (err error) {
	var b []byte
	b, err = ioutil.ReadFile(conn.dataSource + "grbl.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &conn.Controllers)
	if err != nil {
		return
	}
	return
}

func (conn *Connector) Add(gctl *Controller) {
	conn.Controllers[gctl.ID] = gctl
}
