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
	Title        string
	Ports        []string
	Controllers  []Controller
	Controller   Controller
	Views        []string
	View         string
	Theme        string
	Themes       []string
	Icons        []string
	Icon         string
	PowerOn      string
	PowerOff     string
	PowerButtons []string
	Documents    docs.Docs

	// layout *template.Template
	// templates               map[string]*template.Template
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
		Title:        "Evad",
		Ports:        ports,
		Controllers:  controllers,
		Controller:   controllers[0],
		Views:        views,
		View:         views[0],
		Themes:       themes,
		Theme:        "black",
		Icons:        icons,
		Icon:         "white",
		PowerButtons: power_buttons,
		PowerOn:      "green",
		PowerOff:     "gray",
		Documents:    documents,
	}

	Documents = documents

	wp.addRoutes(router, layout)
	// wp.refreshLayout()
	return
}

// func (wp *WebPage) refreshLayout() {
// 	wp.layout = template.Must(template.ParseFiles("../assets/layout.html"))
// 	wp.heading = wp.layout.Lookup("heading")
// 	wp.dialog = wp.layout.Lookup("dialog")
// 	for _, v := range wp.Views {
// 		wp.templates[v] = wp.layout.Lookup(v)
// 	}
// }

// func (wp *WebPage) Refresh(w io.Writer) {
// 	wp.refreshLayout()
// 	wp.layout.Execute(w, wp)
// }

// func (wp *WebPage) WriteView(w io.Writer, view string) {
// 	v := wp.layout.Lookup(view)
// 	v.Execute(w, wp)
// }

// func (wp *WebPage) WriteHeading(w io.Writer) {
// 	wp.heading.Execute(w, wp)
// }

// func (wp *WebPage) WriteDialog(w io.Writer, key string) {
// 	entry := wp.Controller.Item(wp.View, key)
// 	wp.dialog.Execute(w, entry)
// }

var icons = []string{
	"black",
	"blue",
	"gray",
	"green",
	"red",
	"white",
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

var power_buttons = []string{
	"black",
	"blue",
	"gray",
	"green",
	"red",
	"white",
}
