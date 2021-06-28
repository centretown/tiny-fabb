package scrape

import (
	"fmt"
	"os"
	"testing"
	"text/template"

	"github.com/centretown/tiny-fabb/docs"
)

func testGrblScrape(t *testing.T) {
	urls := []string{
		"https://github.com/gnea/grbl/wiki/Grbl-v1.1-Configuration/",
		"https://github.com/gnea/grbl/wiki/Grbl-v1.1-Commands/",
	}
	dsw := docs.NewDocs()
	err := Scrape("h4", urls, dsw, ScrapeGrblWiki)
	if err != nil {
		t.Fatal(err)
	}

	err = dsw.WriteFile("grbl_test.json")
	if err != nil {
		t.Fatal(err)
	}

	dsr := docs.NewDocs()
	err = dsr.ReadFile("grbl_test.json")
	if err != nil {
		t.Fatal(err)
	}

	for _, d := range dsw {
		dr, err := dsr.Find(d.Link)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(dr.Sprint())
	}

	t.Log(dsr.Sprint())
}

func testRepRapScrape(t *testing.T) {
	urls := []string{
		"https://reprap.org/wiki/G-code",
	}

	dsw := docs.NewDocs()
	err := Scrape("#G-commands", urls, dsw, ScrapeRepRapWiki)
	if err != nil {
		t.Fatal(err)
	}
	// err = Scrape("#M-commands", urls, dsw, ScrapeRepRapWiki)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	err = dsw.WriteFile("grbl_test.json")
	if err != nil {
		t.Fatal(err)
	}
}

func testScrape(t *testing.T) {
	documents := docs.NewDocs()

	urlsGrbl := []string{
		"https://github.com/gnea/grbl/wiki/Grbl-v1.1-Configuration/",
		"https://github.com/gnea/grbl/wiki/Grbl-v1.1-Commands/",
	}
	err := Scrape("h4", urlsGrbl, documents, ScrapeGrblWiki)
	if err != nil {
		t.Fatal(err)
	}

	urlsRepRap := []string{
		"https://reprap.org/wiki/G-code",
	}
	err = Scrape("#G-commands", urlsRepRap, documents, ScrapeRepRapWiki)
	if err != nil {
		t.Fatal(err)
	}
	err = Scrape("#M-commands", urlsRepRap, documents, ScrapeRepRapWiki)
	if err != nil {
		t.Fatal(err)
	}
	err = documents.WriteFile("grbl_test.json")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRead(t *testing.T) {
	layout := template.Must(template.ParseFiles("../assets/layout.test"))
	documents := docs.NewDocs()
	documents.ReadFile("grbl_test.json")

	var findAndLayout = func(key string) {
		doc, err := documents.Find(key)
		if err != nil {
			t.Fatal(err)
		}
		err = layout.ExecuteTemplate(os.Stdout, "pop-doc", doc)
		if err != nil {
			t.Fatal(err)
		}
	}

	findAndLayout("G0")
	findAndLayout("G1")
	findAndLayout("G99")
	findAndLayout("G4")
	findAndLayout("M101")
	findAndLayout("$$")
	findAndLayout("$111")
	findAndLayout("G55")
	findAndLayout("G3")
	findAndLayout("G28")
	findAndLayout("G29")

}

func testSplitKeys(t *testing.T) {
	seps := []string{" & ", " and ", ",", "\u0026"}
	rsep := ".."
	testSplit(t, "G0 & G1", seps, rsep, 2)
	testSplit(t, "$110, $111 and $112", seps, rsep, 3)
	testSplit(t, "G17..19", seps, rsep, 3)
	testSplit(t, "G2 \u0026 G3", seps, rsep, 2)
	//"G53..59"
	testSplit(t, "G53..59", seps, rsep, 7)
	testSplit(t, "G47", seps, rsep, 1)
	testSplit(t, "$$", seps, rsep, 1)
	testSplit(t, "$#", seps, rsep, 1)
}

func testSplit(t *testing.T, src string, seps []string, rsep string, count int) {
	t.Log(src)
	keys := splitKeysOrRange(src, seps, rsep)
	if len(keys) != count {
		t.Fatalf("splitKeys count mismatch %q, expected %d, got %d", keys, count, len(keys))
	}
	t.Logf("%d:%v", len(keys), showKeys(keys))
}

func showKeys(keys []string) (s string) {
	for _, k := range keys {
		s += fmt.Sprintf("'%s', ", k)
	}
	return
}
