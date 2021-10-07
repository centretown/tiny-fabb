package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"

	"github.com/centretown/tiny-fabb/forms"
	"github.com/centretown/tiny-fabb/settings"
	"github.com/centretown/tiny-fabb/theme"
	"github.com/centretown/tiny-fabb/web"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

func main() {
	flag.Parse()

	profile := settings.CurrentProfile
	if settings.LoadErr != nil {
		err := profile.Save()
		if err != nil {
			glog.Warningln(err)
		}
	}

	r := profile.Print()
	for _, s := range r {
		glog.Infoln(s)
	}

	router := mux.NewRouter()
	// serve static files from assets folder
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir(profile.AssetsPath+"/"))))

	controllers, ports, layout, documents, camConn := profile.Setup()

	forms.UseDocuments(documents)

	themes := make(theme.Themes)
	err := themes.ReadJSON(profile.DataSource + "/themes.json")
	if err != nil {
		glog.Fatal(err)
	}

	webPage, err := web.NewPage(router, controllers, ports, layout, themes)
	if err != nil {
		glog.Fatal(err)
	}
	webPage.Cameras = camConn.Cameras

	camConn.Start(router, profile.ServoController, profile.Cameras...)
	server := &http.Server{Addr: profile.WebPort, Handler: router}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	go func() {
		<-sc
		camConn.Stop()
		server.Shutdown(context.Background())
	}()

	glog.Infof("Web Server:%s Active", webPage.Title)
	err = server.ListenAndServe()
	glog.Info(err)
	glog.Exit()
}
