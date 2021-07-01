package docs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Doc struct {
	Refer   string            `json:"refer"`
	Title   string            `json:"title"`
	Link    string            `json:"link"`
	Href    string            `json:"href"`
	Text    []string          `json:"text"`
	Support map[string]string `json:"support"`
	Subs    []*Doc            `json:"subs"`
}

func NewDoc() (doc *Doc) {
	doc = &Doc{}
	return
}

type Docs map[string]*Doc

func NewDocs() (docs Docs) {
	docs = make(Docs)
	return
}

func (docs Docs) WriteFile(fileName string) (err error) {
	b, err := json.MarshalIndent(docs, "", "  ")
	if err == nil {
		err = ioutil.WriteFile(fileName, b, 0640)
	}
	return
}

func (docs Docs) ReadFile(fileName string) (err error) {
	var b []byte
	b, err = ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &docs)
	return
}

func (doc *Doc) Sprint() (s string) {
	doc.Title = strings.TrimSpace(doc.Title)
	s = fmt.Sprintf("Title: %s\nLink: %s\n", doc.Title, doc.Link)
	for _, t := range doc.Text {
		s += fmt.Sprintln(t)
	}
	for _, sub := range doc.Subs {
		s += sub.Sprint()
	}

	return
}

func (docs Docs) Sprint() (s string) {
	for k, d := range docs {
		s += fmt.Sprintf("%s: %s\n", k, d.Title)
	}
	return
}

func (docs Docs) Find(key string) (doc *Doc, err error) {
	doc, ok := docs[key]
	if !ok {
		err = fmt.Errorf("Docs key '%s' not found", key)
	}
	if len(doc.Refer) > 0 {
		return docs.Find(doc.Refer)
	}
	return
}

func LoadDocuments(path string) (documents Docs, err error) {
	documents = NewDocs()
	err = documents.ReadFile(path)
	return
}
