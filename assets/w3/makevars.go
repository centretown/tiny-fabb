package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/glog"
)

type Color struct {
	Name            string `json:"name"`
	Color           string `json:"color"`
	BackGroundColor string `json:"backGroundColor"`
	Style           string `json:"style"`
}

type Theme []*Color

type Themes map[string]Theme

func (themes Themes) writeJSON() {
	b, err := json.MarshalIndent(themes, "", "  ")
	if err != nil {
		glog.Fatal(err)
	}

	err = ioutil.WriteFile("vars/themes.json", b, 0642)
	if err != nil {
		glog.Fatal(err)
	}
}

func (theme Theme) makeCSS() (s string) {
	s = ":root {\n"
	for _, col := range theme {
		s += fmt.Sprintf("--%s: %s;\n", col.Name, col.Color)
		if len(col.BackGroundColor) > 0 {
			s += fmt.Sprintf("--%s-bg: %s;\n", col.Name, col.BackGroundColor)
		}
	}
	s += "}\n"
	return
}

func (theme Theme) writeCSS(fname string) {
	fmt.Println(fname)
	s := theme.makeCSS()

	err := ioutil.WriteFile("vars/"+fname, []byte(s), 0642)
	if err != nil {
		glog.Fatal(err)
	}
}

func main() {
	flag.Parse()
	files, err := os.ReadDir(".")
	if err != nil {
		glog.Fatal(err)
	}

	themes := make(Themes)

	for _, f := range files {
		fname := f.Name()
		if strings.HasPrefix(fname, "w3-") &&
			strings.HasSuffix(fname, ".css") {
			theme := makeColors(fname)
			theme.writeCSS(fname)
			themes[fname] = theme
		}
	}

	themes.writeJSON()
}

func makeColors(fname string) (theme Theme) {

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		glog.Fatal(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		color := &Color{}
		if len(line) < 1 {
			continue
		}

		if strings.HasPrefix(line, ".w3-text") || strings.HasPrefix(line, ".w3-hover-text") {
			_, err = fmt.Sscanf(line, ".%s {color:%s !important}",
				&color.Style, &color.Color)
			if err != nil {
				glog.Infoln(err)
			}
		} else if strings.HasPrefix(line, ".w3-border") || strings.HasPrefix(line, ".w3-hover-border") {
			_, err = fmt.Sscanf(line, ".%s {border-color:%s !important}",
				&color.Style, &color.Color)
			if err != nil {
				glog.Infoln(err)
			}
		} else {
			_, err = fmt.Sscanf(line, ".%s {color:%s !important; background-color:%s !important}",
				&color.Style, &color.Color, &color.BackGroundColor)
			if err != nil {
				glog.Infoln(err)
			}
		}

		color.Name = strings.ReplaceAll(color.Style, ":hover", "")
		color.Name = strings.ReplaceAll(color.Name, "-theme", "")
		theme = append(theme, color)
	}
	return
}
