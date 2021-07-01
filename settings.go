package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

type SettingsHelp struct {
	AssetsPath      string `json:"assetsPath"`
	DataSource      string `json:"dataSource"`
	DocsSource      string `json:"docsSource"`
	ControllerCount string `json:"controllerCount"`
	WebPort         string `json:"webPort"`
	UDPPort         string `json:"udpPort"`
}

type Settings struct {
	Help            SettingsHelp `json:"help"`
	AssetsPath      string       `json:"assetsPath"`
	DataSource      string       `json:"dataSource"`
	DocsSource      string       `json:"docsSource"`
	ControllerCount int          `json:"controllerCount"`
	WebPort         string       `json:"webPort"`
	UDPPort         string       `json:"udpPort"`
}

var LocalSettings = &Settings{
	Help: SettingsHelp{
		DataSource:      "location of data folder",
		DocsSource:      "location of documentation file",
		WebPort:         "web server port",
		UDPPort:         "udp port",
		ControllerCount: "number of controllers to allocate",
		AssetsPath:      "location of static assets",
	},
	AssetsPath:      "assets",
	DataSource:      "assets/data/",
	DocsSource:      "assets/data/docs.json",
	WebPort:         ":8080",
	UDPPort:         ":44444",
	ControllerCount: 5,
}

var (
	LoadSettingsErr  error
	SaveSettingsErr  error
	SettingsDirName  = ".tiny_fabb"
	SettingsFileName = "settings.json"
)

func init() {
	LoadSettingsErr = LocalSettings.Load()
	flag.StringVar(&LocalSettings.DataSource, "datasource",
		LocalSettings.DataSource, LocalSettings.Help.DataSource)
	flag.StringVar(&LocalSettings.DocsSource, "docsource",
		LocalSettings.DocsSource, LocalSettings.Help.DocsSource)
	flag.StringVar(&LocalSettings.WebPort, "webport",
		LocalSettings.WebPort, LocalSettings.Help.WebPort)
	flag.StringVar(&LocalSettings.UDPPort, "udpport",
		LocalSettings.UDPPort, LocalSettings.Help.UDPPort)
	flag.IntVar(&LocalSettings.ControllerCount, "count",
		LocalSettings.ControllerCount,
		LocalSettings.Help.ControllerCount)
	flag.StringVar(&LocalSettings.AssetsPath, "assets",
		LocalSettings.AssetsPath, LocalSettings.Help.AssetsPath)
}

func (s *Settings) Save() (err error) {
	byts, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(DefaultSettings(), byts, 0640)
	return
}

func (s *Settings) Load() (err error) {
	byts, err := ioutil.ReadFile(DefaultSettings())
	if err != nil {
		return
	}
	err = json.Unmarshal(byts, s)
	return
}

func DefaultSettings() string {
	usr, _ := user.Current()

	settingsDir := path.Join(usr.HomeDir, SettingsDirName)
	_, err := os.Stat(settingsDir)
	if err != nil {
		SaveSettingsErr = os.Mkdir(settingsDir, 0755)
	}
	return path.Join(settingsDir, SettingsFileName)
}
