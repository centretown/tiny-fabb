package web

import (
	"fmt"
	"html/template"
	"io"

	"github.com/centretown/tiny-fabb/docs"
	"github.com/gorilla/mux"
)

var Documents docs.Docs

func findDoc(code string) (doc *docs.Doc) {
	doc, err := Documents.Find(code)
	if err != nil {
		doc = &docs.Doc{}
	}
	return
}

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
	Apply(view, key string, vals map[string][]string) ([]*Updated, error)
	Query(view, key string) (err error)
}

type Page struct {
	Title       string
	Ports       []string
	Controllers []Controller
	Controller  Controller
	Views       []string
	View        string
	Theme       string
	Themes      []string
	Documents   docs.Docs
}

func NewPage(router *mux.Router,
	controllers []Controller,
	ports []string,
	layout *template.Template,
	documents docs.Docs) (wp *Page, err error) {
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
		Title:       "Evad",
		Ports:       ports,
		Controllers: controllers,
		Controller:  controllers[0],
		Views:       views,
		View:        views[0],
		Themes:      themes,
		Theme:       "blue-grey",
		Documents:   documents,
	}

	Documents = documents

	wp.addRoutes(router, layout)
	return
}

var themes = []string{
	"red",
	"deep-purple",
	"indigo",
	"blue",
	"light-blue",
	"cyan",
	"green",
	"khaki",
	"amber",
	"orange",
	"blue-grey",
	"brown",
	"grey",
	"dark-grey",
	"black",
}
