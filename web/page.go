package web

import (
	"fmt"
	"html/template"

	"github.com/centretown/tiny-fabb/camera"
	"github.com/centretown/tiny-fabb/monitor"
	"github.com/centretown/tiny-fabb/theme"
	"github.com/gorilla/mux"
)

type Page struct {
	Title       string
	Ports       []string
	Controllers []monitor.Controller
	Views       []*monitor.View
	Color       string
	Theme       theme.Theme
	Themes      theme.Themes
	Cameras     camera.Cameras
}

func NewPage(router *mux.Router,
	controllers []monitor.Controller,
	ports []string,
	layout *template.Template,
	themes theme.Themes) (wp *Page, err error) {

	if len(controllers) < 1 {
		err = fmt.Errorf("no controllers")
		return
	}

	controller := controllers[0]
	views := controller.Views()
	if len(views) < 1 {
		err = fmt.Errorf("no views")
		return
	}

	wp = &Page{
		Title:       "Tiny Fabb",
		Ports:       ports,
		Controllers: controllers,
		Views:       views,
		Color:       "blue-grey",
		Themes:      themes,
	}

	ok := false
	wp.Theme, ok = wp.Themes[wp.Color]
	if !ok {
		err = fmt.Errorf("theme not found %s ", wp.Color)
		return
	}

	wp.addRoutes(router, layout)
	return
}
