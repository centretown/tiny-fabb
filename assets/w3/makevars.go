package main

import (
	"flag"
	"os"

	"github.com/centretown/tiny-fabb/theme"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	files, err := os.ReadDir(".")
	if err != nil {
		glog.Fatal(err)
	}

	themes := make(theme.Themes)
	fname := ""

	for _, f := range files {
		fname = f.Name()
		if theme.Filter(fname, "w3-", ".css") {
			theam, err := theme.NewThemeFromFile(fname)
			if err != nil {
				glog.Error(err)
				continue
			}

			err = theam.WriteCSS(fname)
			if err != nil {
				glog.Error(err)
				continue
			}

			glog.Infof("processed %s\n", fname)

			themes[theme.ColorName(fname)] = theam
		}
	}

	err = themes.WriteJSON("../data/themes.json")
	if err != nil {
		glog.Error(err)
		return
	}

	glog.Infoln("processed themes successfully")
}
