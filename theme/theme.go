package theme

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (themes Themes) WriteJSON(fname string) (err error) {
	b, err := json.MarshalIndent(themes, "", "  ")
	if err != nil {
		glog.Error(err)
		return
	}

	err = ioutil.WriteFile(fname, b, 0640)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}

func (themes Themes) ReadJSON(fname string) (err error) {
	var b []byte
	b, err = ioutil.ReadFile(fname)
	if err != nil {
		glog.Error(err)
		return
	}

	err = json.Unmarshal(b, &themes)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}

var themeCSS = `
.w3-theme-l5 {color:var(--w3-l5) !important; background-color:var(--w3-l5-bg) !important}
.w3-theme-l4 {color:var(--w3-l4) !important; background-color:var(--w3-l4-bg) !important}
.w3-theme-l3 {color:var(--w3-l3) !important; background-color:var(--w3-l3-bg) !important}
.w3-theme-l2 {color:var(--w3-l2) !important; background-color:var(--w3-l2-bg) !important}
.w3-theme-l1 {color:var(--w3-l1) !important; background-color:var(--w3-l1-bg) !important}
.w3-theme-d1 {color:var(--w3-d1) !important; background-color:var(--w3-d1-bg) !important}
.w3-theme-d2 {color:var(--w3-d2) !important; background-color:var(--w3-d2-bg) !important}
.w3-theme-d3 {color:var(--w3-d3) !important; background-color:var(--w3-d3-bg) !important}
.w3-theme-d4 {color:var(--w3-d4) !important; background-color:var(--w3-d4-bg) !important}
.w3-theme-d5 {color:var(--w3-d5) !important; background-color:var(--w3-d5-bg) !important}

.w3-theme-light {color:var(--w3-light) !important; background-color:var(--w3-light-bg) !important}
.w3-theme-dark {color:var(--w3-dark) !important; background-color:var(--w3-dark-bg) !important}
.w3-theme-action {color:var(--w3-action) !important; background-color:var(--w3-action-bg) !important}

.w3-theme {color:var(--w3) !important; background-color:var(--w3-bg) !important}
.w3-text-theme {color:var(--w3-text) !important}

.w3-border-theme {border-color:var(--w3-border) !important}
.w3-hover-theme:hover {color:var(--w3-hover) !important; background-color:var(--w3-hover-bg) !important}
.w3-hover-text-theme:hover {color:var(--w3-hover-text) !important}
.w3-hover-border-theme:hover {border-color:var(--w3-hover-border) !important}
`

func MakeThemeCSS() string {
	return themeCSS
}

func (theme Theme) MakeCSS() (s string) {
	s = ":root {\n"
	for _, item := range theme {
		s += fmt.Sprintf("  --%s: %s;\n", item.Name, item.Color)
		if len(item.BackGroundColor) > 0 {
			s += fmt.Sprintf("  --%s-bg: %s;\n",
				item.Name, item.BackGroundColor)
		}
	}
	s += "}\n"
	return
}

func (theme Theme) WriteCSS(fname string) (err error) {
	s := theme.MakeCSS()

	err = ioutil.WriteFile("vars/"+fname, []byte(s), 0640)
	if err != nil {
		glog.Error(err)
		return
	}
	return
}

func ColorName(fname string) (colorName string) {
	_, err := fmt.Sscanf(fname, "w3-theme-%s", &colorName)
	if err != nil {
		glog.Warning(err, fname)
		colorName = fname
		return
	}
	idx := strings.Index(colorName, ".css")
	if idx != -1 {
		colorName = colorName[:idx]
	}
	return
}

func Filter(fname, prefix, suffix string) (result bool) {
	result = strings.HasPrefix(fname, prefix) &&
		strings.HasSuffix(fname, suffix)
	return
}

func isPrefixed(src string, prefixes ...string) (result bool) {
	for _, prefix := range prefixes {
		if strings.HasPrefix(src, prefix) {
			result = true
			return
		}
	}
	return
}

func NewThemeFromFile(fname string) (theme Theme, err error) {

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		glog.Error(err)
		return
	}

	for _, line := range strings.Split(string(b), "\n") {
		if len(line) < 1 {
			continue
		}

		color := &Color{}
		switch {
		case isPrefixed(line, ".w3-text", ".w3-hover-text"):
			_, err = fmt.Sscanf(line,
				".%s {color:%s !important}",
				&color.Style, &color.Color)

		case isPrefixed(line, ".w3-border", ".w3-hover-border"):
			_, err = fmt.Sscanf(line,
				".%s {border-color:%s !important}",
				&color.Style, &color.Color)

		default:
			_, err = fmt.Sscanf(line,
				".%s {color:%s !important; background-color:var%s !important}",
				&color.Style, &color.Color, &color.BackGroundColor)
		}

		if err != nil {
			err = fmt.Errorf("%v: processing %v", err, line)
			return
		}

		color.Name = strings.ReplaceAll(color.Style, ":hover", "")
		color.Name = strings.ReplaceAll(color.Name, "-theme", "")
		theme = append(theme, color)
	}
	return
}
