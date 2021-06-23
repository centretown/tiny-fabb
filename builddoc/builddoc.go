package main

import (
	"flag"

	"github.com/centretown/tiny-fabb/docs"
	"github.com/centretown/tiny-fabb/scrape"
	"github.com/golang/glog"
)

var (
	Destination     string = "../assets/data/docs.json"
	DestinationHelp string = "location of 'docs' data assets"
)

func init() {
	flag.StringVar(&Destination, "datasource",
		Destination, DestinationHelp)
}

func main() {
	flag.Parse()
	flag.VisitAll(func(f *flag.Flag) {
		glog.Infof("%s: %v\n", f.Name, f.Value)
	})

	documents := docs.NewDocs()

	urlsGrbl := []string{
		"https://github.com/gnea/grbl/wiki/Grbl-v1.1-Configuration/",
		"https://github.com/gnea/grbl/wiki/Grbl-v1.1-Commands/",
	}
	err := scrape.Scrape("h4", urlsGrbl, documents, scrape.ScrapeGrblWiki)
	if err != nil {
		glog.Error(err)
		return
	}

	urlsRepRap := []string{
		"https://reprap.org/wiki/G-code",
	}

	err = scrape.Scrape("#G-commands", urlsRepRap, documents, scrape.ScrapeRepRapWiki)
	if err != nil {
		glog.Fatal(err)
	}

	err = scrape.Scrape("#M-commands", urlsRepRap, documents, scrape.ScrapeRepRapWiki)
	if err != nil {
		glog.Fatal(err)
	}
	err = documents.WriteFile(Destination)
	if err != nil {
		glog.Fatal(err)
	}

	glog.Infof("Documents have been successfully scraped and stored to %s.\n", Destination)
}
