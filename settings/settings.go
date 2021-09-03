package settings

import (
	"encoding/json"
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"os/user"
	"path"

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
	Include         string `json:"include"`
	Exclude         string `json:"exclude"`
}

type Profile struct {
	Help            Help     `json:"help"`
	AssetsPath      string   `json:"assetsPath"`
	DataSource      string   `json:"dataSource"`
	DocsSource      string   `json:"docsSource"`
	ControllerCount int      `json:"controllerCount"`
	WebPort         string   `json:"webPort"`
	UDPPort         string   `json:"udpPort"`
	Cameras         []string `json:"cameras"`
	Include         []string `json:"include"`
	Exclude         []string `json:"exclude"`
}

var CurrentProfile = &Profile{
	Help: Help{
		DataSource:      "location of data folder",
		DocsSource:      "location of documentation file",
		WebPort:         "web server port",
		UDPPort:         "udp port",
		ControllerCount: "number of controllers to allocate",
		AssetsPath:      "location of static assets",
		Include:         "device prefixes to include eg: ttyUSB,ttyACM",
		Exclude:         "devices to exclude eg: ttyUSB0",
	},
	AssetsPath:      "assets",
	DataSource:      "assets/data/",
	DocsSource:      "assets/data/docs.json",
	WebPort:         ":8080",
	UDPPort:         ":44444",
	ControllerCount: 5,
	Cameras:         []string{""},
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
	flag.IntVar(&CurrentProfile.ControllerCount, "count",
		CurrentProfile.ControllerCount,
		CurrentProfile.Help.ControllerCount)
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

func (s *Profile) Setup() (controllers []monitor.Controller, ports []string, layout *template.Template, documents docs.Docs) {
	layout = template.Must(
		template.ParseFiles(
			s.AssetsPath+"/layout.html",
			s.AssetsPath+"/layout.go.tpl",
			s.AssetsPath+"/entry.go.tpl"))

	prv := &serialio.Provider{}
	prv.Filter = s.Include
	prv.Exclude = s.Exclude
	connector := &grbl.Connector{}

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

		gctl, err := connector.Connect(bus, layout)
		if err != nil {
			glog.Warning(err)
			continue
		}

		controllers = append(controllers, gctl)
	}

	documents, err := docs.LoadDocuments(s.DocsSource)
	if err != nil {
		glog.Errorln(err)
	}
	return
}
