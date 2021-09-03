package grbl

import (
	"fmt"
	"html/template"

	"github.com/centretown/tiny-fabb/monitor"
)

type Connector struct{}

var controllerCount int = 0

func (conn *Connector) Connect(bus *monitor.Bus, layout *template.Template) (ctl monitor.Controller, err error) {
	var (
		terminators = []string{"ok", "error"}
		results     []string
	)

	results, err = bus.Capture("$I", terminators...)
	if err != nil {
		return
	}

	if len(results) == 0 {
		err = fmt.Errorf("no build information")
		return
	}

	gctl := NewController(bus, layout)

	controllerCount++
	gctl.Title = fmt.Sprintf("Controller-%d", controllerCount)

	ctl = gctl

	err = ctl.Query("settings", idSettings.String())
	if err != nil {
		return
	}

	err = ctl.Query("commands", idParameters.String())
	if err != nil {
		return
	}

	err = ctl.Query("commands", idParserState.String())
	if err != nil {
		return
	}

	err = ctl.Query("commands", idBuildInfo.String())
	if err != nil {
		return
	}

	err = ctl.Query("commands", idStartupBlocks.String())
	if err != nil {
		return
	}

	// err = gctl.Query("commands", idCodeMode.String())
	// if err != nil {
	// 	return
	// }

	return
}
