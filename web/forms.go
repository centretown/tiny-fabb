package web

import "github.com/centretown/tiny-fabb/docs"

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
	return Documents[code]
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

func (form *Form) Update(vals map[string][]string) (err error) {
	for _, entry := range form.Entries {
		ss, ok := vals[entry.ID]
		if ok && len(ss) > 0 {
			err = entry.ScanInput(ss[0], form.Value)
		} else if entry.Type == "checkbox" {
			// no values are submitted for unchecked checkboxes
			err = entry.ScanInput("false", form.Value)
		}
		if err != nil {
			return
		}
	}
	return
}
