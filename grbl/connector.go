package grbl

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"

	"github.com/centretown/tiny-fabb/monitor"
)

type Connector struct {
	List        []*Controller
	Controllers map[string]*Controller
	dataSource  string
	layout      *template.Template
}

func NewConnector(dataSource string, layout *template.Template) (conn *Connector) {
	conn = &Connector{
		List:        make([]*Controller, 0),
		Controllers: make(map[string]*Controller),
		dataSource:  dataSource,
		layout:      layout,
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
		controllerCount int = len(conn.List)
	)

	results, err = bus.Capture("$I", terminators...)
	if err != nil {
		return
	}

	profile, err = CheckProfile(results)
	if err != nil {
		return
	}

	if len(profile.ID) == 0 {
		profile.ID = fmt.Sprintf("%s-%03d", profile.Prefix(), controllerCount+1)
		s := fmt.Sprintf("$I=%s%s", ID, profile.ID)
		_, err = bus.Capture(s, terminators...)
		if err != nil {
			return
		}
	}

	gctl, ok = conn.Controllers[profile.ID]
	if ok {
		gctl.Profile = profile
		gctl.initialize(bus, conn.layout)
	} else {
		controllerCount++
		gctl = NewController(bus, conn.layout)
		gctl.Profile = profile
		gctl.ID = profile.ID
		gctl.Title = fmt.Sprintf("%s-%03d", profile.Prefix(), controllerCount)
		conn.Add(gctl)
	}

	ctl = gctl
	err = gctl.startup()
	return
}

func (conn *Connector) Save() (err error) {
	var b []byte
	b, err = json.MarshalIndent(&conn.List, "", "  ")
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
	err = json.Unmarshal(b, &conn.List)
	if err != nil {
		return
	}
	for _, ctl := range conn.List {
		conn.Controllers[ctl.ID] = ctl
	}
	return
}

func (conn *Connector) Add(gctl *Controller) {
	conn.List = append(conn.List, gctl)
	conn.Controllers[gctl.ID] = gctl
}
