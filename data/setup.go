package data

import (
	"html/template"

	"github.com/centretown/tiny-fabb/docs"
	"github.com/centretown/tiny-fabb/grbl"
	"github.com/centretown/tiny-fabb/monitor"
	"github.com/centretown/tiny-fabb/serialio"
	"github.com/golang/glog"
)

func Setup(assets, docsSource string) (controllers []monitor.Controller, ports []string, layout *template.Template, documents docs.Docs) {
	layout = template.Must(
		template.ParseFiles(
			assets+"/layout.html",
			assets+"/layout.go.tpl",
			assets+"/entry.go.tpl"))

	prv := &serialio.Provider{}
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

	documents, err := docs.LoadDocuments(docsSource)
	if err != nil {
		glog.Errorln(err)
	}
	return
}
