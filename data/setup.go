package data

import (
	"fmt"
	"html/template"

	"github.com/centretown/tiny-fabb/grbl"
	"github.com/centretown/tiny-fabb/serialio"
	"github.com/centretown/tiny-fabb/web"
	"github.com/golang/glog"
)

func Setup(count int, assets string) (controllers []web.Controller, ports []string, layout *template.Template) {
	layout = template.Must(template.ParseFiles(assets + "/layout.html"))
	ports = serialio.ListSerial()
	glog.Infoln(ports)
	controllers = make([]web.Controller, count)
	for i, _ := range controllers {
		g := grbl.NewController(layout)
		g.Title = fmt.Sprintf("GRBL-%02d", i+1)

		if i < len(ports) {
			g.Port = ports[i]
			err := g.ActivateController()
			if err != nil {
				glog.Warningln(err)
			}
		}
		controllers[i] = g
	}
	return
}
