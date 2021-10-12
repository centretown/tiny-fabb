// Copyright (c) 2021 Dave Marsh. See LICENSE.

package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/centretown/tiny-fabb/forms"
	"github.com/centretown/tiny-fabb/monitor"
	"github.com/gorilla/mux"
)

const (
	selController = "controller"
	selView       = "view"
	selKey        = "key"
	urlView       = "{controller:[0-9]+}/{view}/"
	urlKey        = urlView + "{key}/"
)

func (wp *Page) addRoutes(router *mux.Router, layout *template.Template) {
	router.HandleFunc("/view/"+urlView, wp.handleView)
	router.HandleFunc("/edit/"+urlKey, wp.handleEdit)
	router.HandleFunc("/apply/"+urlKey, wp.handleApply)
	router.HandleFunc("/options/{theme}/", wp.handleOptions)
	router.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			layout.Execute(w, wp)
		})
}

func (wp *Page) handleView(w http.ResponseWriter, r *http.Request) {
	controller, view, err := wp.GetView(w, r)
	if err != nil {
		return
	}
	err = controller.View(w, view)
	if err != nil {
		forms.WriteError(w, err)
		return
	}
}

func (wp *Page) handleEdit(w http.ResponseWriter, r *http.Request) {
	controller, view, err := wp.GetView(w, r)
	if err != nil {
		return
	}

	key := forms.GetRequestString(r, "key")
	err = controller.Edit(w, view, key)
	if err != nil {
		forms.WriteError(w, err)
		return
	}
}

func (wp *Page) handleApply(w http.ResponseWriter, r *http.Request) {
	controller, view, err := wp.GetView(w, r)
	if err != nil {
		return
	}

	var (
		updated []*forms.Updated
		pkg     []byte
	)
	key := forms.GetRequestString(r, "key")
	err = r.ParseForm()
	if err == nil || err == io.EOF {
		updated, err = controller.Apply(view, key, r.Form)
		if err == nil {
			pkg, err = json.Marshal(updated)
		}
	}

	if err != nil {
		forms.WriteError(w, err)
		return
	}
	fmt.Fprint(w, string(pkg))
}

func (wp *Page) handleOptions(w http.ResponseWriter, r *http.Request) {
	var ok bool
	theme := forms.GetRequestString(r, "theme")
	wp.Theme, ok = wp.Themes[theme]
	if ok {
		s := wp.Theme.MakeCSS()
		w.Write([]byte(s))
	}
}

func (wp *Page) GetView(w http.ResponseWriter, r *http.Request) (ctlr monitor.Controller, view string, err error) {
	sel := forms.GetRequestUint(r, selController)
	if int(sel) >= len(wp.Controllers) {
		err = fmt.Errorf("controller %d not found", sel)
		forms.WriteError(w, err)
		return
	}
	ctlr = wp.Controllers[sel]
	view = forms.GetRequestString(r, selView)
	if !wp.findView(view) {
		err = fmt.Errorf("view %s not found", view)
		forms.WriteError(w, err)
		return
	}
	return
}

func (wp *Page) findView(view string) bool {
	for _, v := range wp.Views {
		if v.ID == view {
			return true
		}
	}
	return false
}
