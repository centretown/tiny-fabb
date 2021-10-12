package forms

import (
	"encoding/json"
	"fmt"

	"github.com/centretown/tiny-fabb/docs"
)

var documents docs.Docs

func UseDocuments(init docs.Docs) {
	documents = init
}

func Documents() docs.Docs {
	return documents
}

func findDoc(code string) (doc *docs.Doc) {
	doc, err := documents.Find(code)
	if err != nil {
		doc = &docs.Doc{}
	}
	return
}

type Form struct {
	ID      WebId
	Value   interface{}
	Entries Entries
	BaseUrl string
}

type Forms map[WebId]*Form

type Updated struct {
	ID    string      `json:"id"`
	Value interface{} `json:"value"`
}

func (form *Form) FindDoc(code string) *docs.Doc {
	return findDoc(code)
}

func (form *Form) GetUpdated() (updated []*Updated) {
	updated = make([]*Updated, 0)
	for _, entry := range form.Entries {
		updated = append(updated,
			&Updated{
				ID:    entry.ID,
				Value: entry.Value(form.Value),
			})
	}
	return
}

type Identifier struct {
	CtlID   string
	View    string
	FormID  string
	Form    *Form
	Icon    string
	Results []string
}

func (form *Form) Identify(ctlid, view, icon string,
	results []string) (ident *Identifier) {

	ident = &Identifier{
		CtlID:   ctlid,
		View:    view,
		FormID:  fmt.Sprintf("%s-%s", ctlid, form.ID.String()),
		Icon:    icon,
		Results: results,
		Form:    form,
	}
	return
}

func (form *Form) Update(vals map[string][]string) (err error) {
	for _, entry := range form.Entries {
		ss, ok := vals[entry.ID]
		if entry.Type == "checkbox" {
			if ok {
				err = entry.ScanInput("true", form.Value)
			} else {
				err = entry.ScanInput("false", form.Value)
			}
		} else if ok && len(ss) > 0 {
			err = entry.ScanInput(ss[0], form.Value)
		}

		if err != nil {
			return
		}
	}
	return
}

func (form *Form) ToJSON() (b []byte) {
	m := make(map[string]interface{})
	for _, entry := range form.Entries {
		m[entry.ID] = entry.Value((form.Value))
	}
	b, _ = json.MarshalIndent(m, "", "  ")
	return
}
