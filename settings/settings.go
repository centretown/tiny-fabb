package settings

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"os/user"
	"path"

	"github.com/centretown/tiny-fabb/camera"
	"github.com/centretown/tiny-fabb/docs"
	"github.com/centretown/tiny-fabb/grbl"
	"github.com/centretown/tiny-fabb/monitor"
	"github.com/centretown/tiny-fabb/serialio"
	"github.com/golang/glog"
)

type Help struct {
	AssetsPath      string `json:"assetsPath"`
	DataSource      string `json:"dataSource"`
	DocsSource      string `json:"docsSource"`
	ControllerCount string `json:"controllerCount"`
	WebPort         string `json:"webPort"`
	UDPPort         string `json:"udpPort"`
	Cameras         string `json:"cameras"`
	ServoController string `json:"servoController"`
	Include         string `json:"include"`
	Exclude         string `json:"exclude"`
}

type Profile struct {
	Help            Help     `json:"help"`
	AssetsPath      string   `json:"assetsPath"`
	DataSource      string   `json:"dataSource"`
	DocsSource      string   `json:"docsSource"`
	WebPort         string   `json:"webPort"`
	UDPPort         string   `json:"udpPort"`
	Cameras         []string `json:"cameras"`
	ServoController string   `json:"servoController"`
	Include         []string `json:"include"`
	Exclude         []string `json:"exclude"`
}

var CurrentProfile = &Profile{
	Help: Help{
		DataSource:      "location of data folder",
		DocsSource:      "location of documentation file",
		WebPort:         "web server port",
		UDPPort:         "udp port",
		Cameras:         "attached camera addresses",
		ServoController: "servo controller address",
		AssetsPath:      "location of static assets",
		Include:         "device prefixes to include",
		Exclude:         "devices to exclude",
	},
	AssetsPath:      "assets",
	DataSource:      "assets/data/",
	DocsSource:      "assets/data/docs.json",
	WebPort:         ":8080",
	UDPPort:         ":44444",
	Cameras:         []string{""},
	ServoController: "http://192.168.0.44",
	Include:         []string{""},
	Exclude:         []string{""},
}

var (
	LoadErr  error
	SaveErr  error
	DirName  = ".tiny_fabb"
	FileName = "settings.json"
)

func init() {
	LoadErr = CurrentProfile.Load()
	flag.StringVar(&CurrentProfile.DataSource, "datasource",
		CurrentProfile.DataSource, CurrentProfile.Help.DataSource)
	flag.StringVar(&CurrentProfile.DocsSource, "docsource",
		CurrentProfile.DocsSource, CurrentProfile.Help.DocsSource)
	flag.StringVar(&CurrentProfile.WebPort, "webport",
		CurrentProfile.WebPort, CurrentProfile.Help.WebPort)
	flag.StringVar(&CurrentProfile.UDPPort, "udpport",
		CurrentProfile.UDPPort, CurrentProfile.Help.UDPPort)
	flag.StringVar(&CurrentProfile.AssetsPath, "assets",
		CurrentProfile.AssetsPath, CurrentProfile.Help.AssetsPath)
}

func (s *Profile) Save() (err error) {
	byts, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(DefaultProfile(), byts, 0640)
	return
}

func (s *Profile) Load() (err error) {
	byts, err := ioutil.ReadFile(DefaultProfile())
	if err != nil {
		return
	}
	err = json.Unmarshal(byts, s)
	return
}

func DefaultProfile() string {
	usr, _ := user.Current()

	settingsDir := path.Join(usr.HomeDir, DirName)
	_, err := os.Stat(settingsDir)
	if err != nil {
		SaveErr = os.Mkdir(settingsDir, 0755)
	}
	return path.Join(settingsDir, FileName)
}

func (s *Profile) Print() (r []string) {
	r = append(r, fmt.Sprintf("%s: %v", s.Help.AssetsPath, s.AssetsPath))
	r = append(r, fmt.Sprintf("%s: %v", s.Help.DataSource, s.DataSource))
	r = append(r, fmt.Sprintf("%s: %v", s.Help.DocsSource, s.DocsSource))
	r = append(r, fmt.Sprintf("%s: %v", s.Help.WebPort, s.WebPort))
	r = append(r, fmt.Sprintf("%s: %v", s.Help.UDPPort, s.UDPPort))
	r = append(r, fmt.Sprintf("%s: %v", s.Help.Cameras, s.Cameras))
	r = append(r, fmt.Sprintf("%s: %v", s.Help.ServoController, s.ServoController))
	r = append(r, fmt.Sprintf("%s: %v", s.Help.Include, s.Include))
	r = append(r, fmt.Sprintf("%s: %v", s.Help.Exclude, s.Exclude))
	return
}

func (s *Profile) Setup() (controllers []monitor.Controller, ports []string, layout *template.Template, documents docs.Docs, camConn *camera.Connector) {
	layout = template.Must(
		layout.ParseFiles(
			s.AssetsPath+"/layout.html",
			s.AssetsPath+"/layout.go.tpl",
			s.AssetsPath+"/entry.go.tpl",
			s.AssetsPath+"/camera.go.tpl",
			s.AssetsPath+"/servo.go.tpl"))

	prv := &serialio.Provider{}
	prv.Filter = s.Include
	prv.Exclude = s.Exclude
	camConn = camera.NewConnector(s.DataSource, layout, 200)
	fmt.Println(camConn.Cameras)
	grblConn := grbl.NewConnector(s.DataSource, layout, camConn.Cameras)

	ports = prv.Update()
	glog.Infoln(ports)

	controllers = make([]monitor.Controller, 0)
	for _, p := range ports {
		sio, err := prv.Get(p)
		if err != nil {
			glog.Warning(err)
			continue
		}

		bus := monitor.NewBus()
		go monitor.Monitor(sio, bus)

		gctl, err := grblConn.Connect(bus)
		if err != nil {
			glog.Warning(err)
			bus.Done <- true
			continue
		}

		controllers = append(controllers, gctl)
	}

	documents, err := docs.LoadDocuments(s.DocsSource)
	if err != nil {
		glog.Errorln(err)
	}

	err = grblConn.Save()
	if err != nil {
		glog.Errorln(err)
	}

	return
}
