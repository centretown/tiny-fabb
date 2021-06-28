package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

const (
	selected = "selected"
)

func (wp *Page) addRoutes(router *mux.Router, layout *template.Template) {

	router.HandleFunc("/controller/{selected:[0-9]+}/",
		func(w http.ResponseWriter, r *http.Request) {
			sel := getRequestUint(r, selected)
			wp.Controller = wp.Controllers[sel]
			layout.Execute(w, wp)
		})

	router.HandleFunc("/options/{theme}/",
		func(w http.ResponseWriter, r *http.Request) {
			wp.Theme = getRequestString(r, "theme")
			layout.Execute(w, wp)
		})

	router.HandleFunc("/list/{view}/",
		func(w http.ResponseWriter,
			r *http.Request) {
			view := getRequestString(r, "view")
			err := wp.Controller.List(w, view)
			if err != nil {
				glog.Infoln(err)
				http.Error(w, err.Error(), 400)
				return
			}
		})

	router.HandleFunc("/edit/{view}/{key}/",
		func(w http.ResponseWriter,
			r *http.Request) {
			view := getRequestString(r, "view")
			key := getRequestString(r, "key")
			err := wp.Controller.Edit(w, view, key)
			if err != nil {
				glog.Warning(err)
				http.Error(w, err.Error(), 400)
				return
			}
		})

	router.HandleFunc("/apply/{key}/",
		func(w http.ResponseWriter, r *http.Request) {
			var (
				updated []*Updated
				pkg     []byte
			)
			key := getRequestString(r, "key")
			err := r.ParseForm()
			if err == nil || err == io.EOF {
				updated, err = wp.Controller.Apply("settings", key, r.Form)
				if err == nil {
					pkg, err = json.Marshal(updated)
					if err != nil {
						glog.Warningln(err)
						http.Error(w, err.Error(), 400)
						return
					}
				}
			}

			if err != nil {
				glog.Warningln(err)
				http.Error(w, err.Error(), 400)
				return
			}
			fmt.Fprint(w, string(pkg))
		})

	router.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			layout.Execute(w, wp)
		})

	// for _, g := range Controllers {
	// 	err = ActivateController(g)
	// 	if err != nil {
	// 		glog.Infof("%v: %v\n", g.Title, err)
	// 		continue
	// 	}
	// }
}
func getRequestUint(r *http.Request, key string) (sel uint) {
	vars := mux.Vars(r)
	fmt.Sscan(vars[key], &sel)
	return
}

func getRequestString(r *http.Request, key string) (sel string) {
	vars := mux.Vars(r)
	fmt.Sscan(vars[key], &sel)
	return
}
