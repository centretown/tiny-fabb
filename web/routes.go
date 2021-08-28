package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/centretown/tiny-fabb/forms"
	"github.com/centretown/tiny-fabb/monitor"
	"github.com/golang/glog"
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
	router.HandleFunc("/list/"+urlView, wp.handleList)
	router.HandleFunc("/edit/"+urlKey, wp.handleEdit)
	router.HandleFunc("/apply/"+urlKey, wp.handleApply)
	router.HandleFunc("/options/{theme}/", wp.handleOptions)
	router.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			layout.Execute(w, wp)
		})
}

func (wp *Page) handleList(w http.ResponseWriter, r *http.Request) {
	controller, view, err := wp.GetView(w, r)
	if err != nil {
		return
	}
	err = controller.List(w, view)
	if err != nil {
		writeError(w, err)
		return
	}
}

func (wp *Page) handleEdit(w http.ResponseWriter, r *http.Request) {
	controller, view, err := wp.GetView(w, r)
	if err != nil {
		return
	}

	key := getRequestString(r, "key")
	err = controller.Edit(w, view, key)
	if err != nil {
		writeError(w, err)
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
	key := getRequestString(r, "key")
	err = r.ParseForm()
	if err == nil || err == io.EOF {
		updated, err = controller.Apply(view, key, r.Form)
		if err == nil {
			pkg, err = json.Marshal(updated)
		}
	}

	if err != nil {
		writeError(w, err)
		return
	}
	fmt.Fprint(w, string(pkg))
}

func (wp *Page) handleOptions(w http.ResponseWriter, r *http.Request) {
	var ok bool
	theme := getRequestString(r, "theme")
	wp.Theme, ok = wp.Themes[theme]
	if ok {
		s := wp.Theme.MakeCSS()
		w.Write([]byte(s))
	}
}

func (wp *Page) GetView(w http.ResponseWriter, r *http.Request) (ctlr monitor.Controller, view string, err error) {
	sel := getRequestUint(r, selController)
	if int(sel) >= len(wp.Controllers) {
		err = fmt.Errorf("controller %d not found", sel)
		writeError(w, err)
		return
	}
	ctlr = wp.Controllers[sel]
	view = getRequestString(r, selView)
	if !wp.findView(view) {
		err = fmt.Errorf("view %s not found", view)
		writeError(w, err)
		return
	}
	return
}

func (wp *Page) findView(view string) bool {
	for _, v := range wp.Views {
		if v == view {
			return true
		}
	}
	return false
}

func getRequestString(r *http.Request, key string) (sel string) {
	vars := mux.Vars(r)
	fmt.Sscan(vars[key], &sel)
	return
}

func getRequestUint(r *http.Request, key string) (sel uint) {
	vars := mux.Vars(r)
	fmt.Sscan(vars[key], &sel)
	return
}

func writeError(w http.ResponseWriter, err error) {
	glog.Infoln(err)
	http.Error(w, err.Error(), 400)
}
