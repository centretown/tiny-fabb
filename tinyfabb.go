package main

import (
	"flag"
	"net/http"

	"github.com/centretown/tiny-fabb/data"
	"github.com/centretown/tiny-fabb/forms"
	"github.com/centretown/tiny-fabb/theme"
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

	router := mux.NewRouter()
	// serve static files from assets folder
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir(LocalSettings.AssetsPath+"/"))))

	controllers, ports, layout, documents := data.Setup(LocalSettings.ControllerCount,
		LocalSettings.AssetsPath, LocalSettings.DocsSource)

	forms.UseDocuments(documents)

	themes := make(theme.Themes)
	err := themes.ReadJSON(LocalSettings.DataSource + "/themes.json")
	if err != nil {
		return
	}

	webPage, err := web.NewPage(router, controllers, ports, layout, themes)
	if err != nil {
		glog.Fatal(err)
	}

	glog.Infof("Web Server:%s Active", webPage.Title)
	err = http.ListenAndServe(LocalSettings.WebPort, router)
	if err != nil {
		glog.Errorln(err)
	}

}
