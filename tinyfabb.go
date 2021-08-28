package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"

	"github.com/centretown/tiny-fabb/camera"
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

	controllers, ports, layout, documents := data.Setup(LocalSettings.AssetsPath, LocalSettings.DocsSource)

	forms.UseDocuments(documents)

	themes := make(theme.Themes)
	err := themes.ReadJSON(LocalSettings.DataSource + "/themes.json")
	if err != nil {
		glog.Fatal(err)
	}

	webPage, err := web.NewPage(router, controllers, ports, layout, themes)
	if err != nil {
		glog.Fatal(err)
	}

	webPage.Cameras = make(camera.Cameras)
	webPage.Cameras.Start(router, 200, LocalSettings.Cameras...)

	server := &http.Server{Addr: LocalSettings.WebPort, Handler: router}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	go func() {
		<-sc
		webPage.Cameras.Stop()
		server.Shutdown(context.Background())
	}()

	glog.Infof("Web Server:%s Active", webPage.Title)
	err = server.ListenAndServe()
	glog.Info(err)
	glog.Exit()
}
