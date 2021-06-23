package main

import (
	"flag"
	"net/http"

	"github.com/centretown/tiny-fabb/data"
	"github.com/centretown/tiny-fabb/web"
	"github.com/golang/glog"

	"github.com/gorilla/mux"
)

func main() {
	flag.Parse()
	flag.VisitAll(func(f *flag.Flag) {
		glog.Infof("%s: %v\n", f.Name, f.Value)
	})

	if LoadSettingsErr != nil {
		err := LocalSettings.Save()
		if err != nil {
			glog.Warningln(err)
		}
	}

	glog.Infoln(DefaultSettings())
	// if !LocalSettings.Exists() {
	// 	err := LocalSettings.Save()
	// 	if err != nil {
	// 		glog.Warningln(err)
	// 	}
	// }

	router := mux.NewRouter()
	// server static files from assets folder
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir(LocalSettings.AssetsPath+"/"))))

	controllers, ports, layout := data.Setup(LocalSettings.ControllerCount,
		LocalSettings.AssetsPath)
	webPage, err := web.NewPage(router, controllers, ports, layout)
	if err != nil {
		glog.Fatal(err)
	}

	glog.Infof("Web Server:%s Active", webPage.Title)
	http.ListenAndServe(LocalSettings.WebPort, router)
}
