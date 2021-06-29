package web

import (
	"fmt"
	"html/template"
	"io"

	"github.com/centretown/tiny-fabb/forms"
	"github.com/centretown/tiny-fabb/theme"
	"github.com/gorilla/mux"
)

type Descriptor struct {
	Title       string
	Port        string
	Version     string
	Active      bool
	Description string
}

type Controller interface {
	Describe() (d *Descriptor)
	ListViews() (vs []string)
	Upload(w io.Writer, files []string) (err error)
	List(w io.Writer, view string) (err error)
	Edit(w io.Writer, view, key string) (err error)
	Apply(view, key string, vals map[string][]string) ([]*forms.Updated, error)
	Query(view, key string) (err error)
}

type Page struct {
	Title       string
	Ports       []string
	Controllers []Controller
	Controller  Controller
	Views       []string
	View        string
	Color       string
	Theme       theme.Theme
	Themes      theme.Themes
}

func NewPage(router *mux.Router,
	controllers []Controller,
	ports []string,
	layout *template.Template,
	themes theme.Themes) (wp *Page, err error) {

	if len(controllers) < 1 {
		err = fmt.Errorf("no controllers")
		return
	}

	controller := controllers[0]
	views := controller.ListViews()
	if len(views) < 1 {
		err = fmt.Errorf("no views")
		return
	}

	wp = &Page{
		Title:       "Tiny Fabb",
		Ports:       ports,
		Controllers: controllers,
		Controller:  controllers[0],
		Views:       views,
		View:        views[0],
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
