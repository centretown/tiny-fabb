package scrape

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/centretown/tiny-fabb/docs"
	"github.com/golang/glog"

	"github.com/gocolly/colly/v2"
)

type Scraper func(docs.Docs) func(*colly.HTMLElement)

var (
	tokenSeparators = []string{" & ", " and ", ",", "\u0026"}
	rangeSeparator  = ".."
)

func Scrape(selection string, urls []string, ds docs.Docs, scraper Scraper) (err error) {
	c := colly.NewCollector()
	c.OnHTML(selection, scraper(ds))
	c.OnRequest(func(r *colly.Request) {
		glog.Infoln("Visiting", r.URL.String())
	})

	for _, url := range urls {
		err = c.Visit(url)
		if err != nil {
			return
		}
	}
	return
}

func ScrapeGrblWiki(ds docs.Docs) (scan func(e *colly.HTMLElement)) {
	scan = func(e *colly.HTMLElement) {
		href := e.ChildAttr("a", "href")
		if len(href) == 0 {
			return
		}
		flds := splitTitle(strings.TrimSpace(e.Text), "–-")
		var (
			title = ""
			link  = strings.TrimSpace(flds[0])
		)
		if len(flds) > 1 {
			title = strings.TrimSpace(flds[1])
		} else {
			title = link
		}
		doc := &docs.Doc{
			Title: title,
			Link:  link,
			Href:  href,
		}

		selection := e.DOM.NextUntil("h4")
		for _, n := range selection.Nodes {
			text := strings.TrimSpace(goquery.NewDocumentFromNode(n).Text())
			if len(text) > 0 {
				for i := strings.Index(text, "\n"); i != -1; {
					if i > 0 {
						doc.Text = append(doc.Text, text[:i])
					}
					text = text[i+1:]
					i = strings.Index(text, "\n")
				}
				if len(text) > 0 {
					doc.Text = append(doc.Text, text)
				}
			}
		}

		links := splitKeysOrRange(doc.Link, tokenSeparators, rangeSeparator)
		linkDoc(ds, doc, links)
	}
	return
}

func ScrapeRepRapWiki(ds docs.Docs) (scan func(e *colly.HTMLElement)) {

	scan = func(e *colly.HTMLElement) {
		h2 := e.DOM.Parent()
		codeSelection := h2.NextUntil("h2")
		docHeading := codeSelection.Filter("h4")
		for _, n := range docHeading.Nodes {
			nodeDoc := goquery.NewDocumentFromNode(n)
			doc := newDocFromNode(nodeDoc)
			// href := n.ChildAttr("a", "href")

			selection := nodeDoc.NextFilteredUntil("div", "dl").Find("th, td")
			doc.Support = support(selection)

			selection = nodeDoc.NextFilteredUntil("p", "dl")
			doc.Text = append(doc.Text,
				strings.FieldsFunc(selection.Text(),
					func(r rune) bool { return r == '\n' })...)

			dl := nodeDoc.NextFilteredUntil("dl", "h4")
			doc.Subs = append(doc.Subs, newDocsFromDataList(dl)...)

			links := splitKeysOrRange(doc.Link, tokenSeparators, rangeSeparator)
			linkDoc(ds, doc, links)
		}

	}
	return
}

func linkDoc(ds docs.Docs, doc *docs.Doc, links []string) {
	for _, link := range links {
		ex, ok := ds[link]
		if !ok {
			ds[link] = doc
		} else {
			ex.Subs = append(ex.Subs, doc)
		}
	}
}

func support(dt *goquery.Selection) (kvs map[string]string) {
	kvs = make(map[string]string)
	keys := strings.Fields(dt.Filter("th").Text())
	vals := dt.Filter("td")

	for i, n := range vals.Nodes {
		j := i + 1
		if j < len(keys) {
			kvs[keys[j]] = strings.TrimSpace(goquery.NewDocumentFromNode(n).Text())
		}
	}

	return
}

func newDocFromNode(nodeDoc *goquery.Document) (ds *docs.Doc) {
	ds = docs.NewDoc()
	text := nodeDoc.Children().Text()
	sep := strings.Index(text, ":")
	if sep < 0 {
		ds.Link = text
		ds.Title = text
	} else {
		ds.Link = text[:sep]
		ds.Title = strings.TrimSpace(text[sep+1:])
	}
	return
}

func newDocsFromDataList(selection *goquery.Selection) (dsl []*docs.Doc) {
	dt := selection.Children().Filter("dt")
	for _, n := range dt.Nodes {
		nodeDoc := goquery.NewDocumentFromNode(n)
		sub := docs.NewDoc()
		sub.Title = nodeDoc.Text()
		dd := nodeDoc.NextFilteredUntil("dd", "dt")
		for _, n := range dd.Nodes {
			tx := goquery.NewDocumentFromNode(n)
			sub.Text = append(sub.Text, tx.Text())
		}
		if len(sub.Text) > 0 {
			dsl = append(dsl, sub)
		}
	}
	return
}

func splitTitle(s string, sep string) (flds []string) {
	flds = strings.FieldsFunc(s,
		func(r rune) (b bool) {
			for _, s := range sep {
				if r == s {
					b = true
					return
				}
			}
			return
		})
	return
}

func splitKeysOrRange(source string, seps []string, rangeSep string) (keys []string) {
	keys = splitKeys(source, seps)
	if len(keys) == 1 {
		keys = splitRange(source, rangeSep)
	}
	return
}

func splitKeys(source string, seps []string) (keys []string) {

	var (
		src = strings.TrimSpace(source)
		f   = func(src, sep string) string {
			return strings.Replace(src, sep, " ", -1)
		}
		found = func() bool {
			for _, sep := range seps {
				i := strings.Index(source, sep)
				if i != -1 {
					return true
				}
			}
			return false
		}
	)

	if !found() {
		keys = append(keys, source)
		return
	}

	for _, sep := range seps {
		src = f(src, sep)
	}

	keys = strings.Fields(src)
	return
}

func splitRange(source string, sep string) (keys []string) {
	i := strings.Index(source, sep)
	if i == -1 {
		keys = append(keys, source)
		return
	}

	code := source[:1]
	prefix := source[1:i]
	suffix := source[i+len(sep):]

	var (
		first, last int
	)

	fmt.Sscanf(prefix, "%d", &first)
	fmt.Sscanf(suffix, "%d", &last)
	if last > first {
		for j := first; j <= last; j++ {
			keys = append(keys,
				fmt.Sprintf("%s%d", code, j))
		}
	}
	return
}
