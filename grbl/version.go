// Copyright (c) 2021 Dave Marsh. See LICENSE.
package grbl

import (
	"fmt"
	"strings"
)

const (
	VER                = "VER:"
	OPT                = "OPT:"
	MAC                = "MAC="
	IP                 = "IP="
	ID                 = "ID="
	WIFI_STATUS        = "Status="
	ESP32_VERSION      = "1.3a.20210424"
	ATMEGA328P_VERSION = "1.1h.20190825"
	ESP32              = "ESP32"
	ATMEGA328P         = "ATMEGA328P"
)

type Version struct {
	Version   string `json:"version"`
	AxisCount int    `json:"axisCount"`
	WiFi      bool   `json:"wifi"`
	Bluetooth bool   `json:"bluetooth"`
}

type Versions map[string]*Version

type Profile struct {
	ID        string `json:"id"`
	Version   string `json:"version"`
	Options   string `json:"options"`
	AxisCount int    `json:"axisCount"`
	IP        string `json:"ip"`
	Status    string `json:"status"`
}

var grblVersions = Versions{
	ATMEGA328P_VERSION: {Version: ATMEGA328P, AxisCount: 3},
	ESP32_VERSION:      {Version: ESP32_VERSION, AxisCount: 6, WiFi: true, Bluetooth: true},
}

func CheckProfile(info []string) (profile *Profile, err error) {
	v := scanInfo(info, VER)
	if len(v) == 0 {
		return
	}

	version, ok := grblVersions[v]
	if !ok {
		err = fmt.Errorf("version '%s' not found", v)
		return
	}

	profile = &Profile{
		Version:   v,
		Options:   scanInfo(info, OPT),
		AxisCount: version.AxisCount,
	}

	if profile.IsESP32() {
		profile.ID = scanInfo(info, MAC)
		profile.IP = scanInfo(info, IP)
		profile.Status = scanInfo(info, WIFI_STATUS)
	} else {
		profile.ID = scanInfo(info, ID)
	}

	return
}

func (pr *Profile) Prefix() string {
	if pr.IsESP32() {
		return ESP32
	}
	return ATMEGA328P
}

func (pr *Profile) IsESP32() bool {
	return strings.Compare(pr.Version, ESP32_VERSION) == 0
}

func (pr *Profile) IsMEGA328P() bool {
	return strings.Compare(pr.Version, ATMEGA328P_VERSION) == 0
}

func (pr *Profile) Print() (s string) {
	s = fmt.Sprintf("id=%s v=%s o=%s c=%d ip=%s %s",
		pr.ID, pr.Version, pr.Options,
		pr.AxisCount, pr.IP, pr.Status)
	return
}

func scanInfo(info []string, key string) (keyval string) {
	var start, end int
	for _, item := range info {
		start = strings.Index(item, key)
		if start == -1 {
			continue
		}

		start += len(key)
		item = item[start:]
		end = strings.IndexAny(item, ":]")
		if end == -1 {
			continue
		}

		keyval = item[:end]
		return
	}
	return
}
