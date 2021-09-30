package monitor

import (
	"io"

	"github.com/centretown/tiny-fabb/forms"
)

type Connector interface {
	Connect(*Bus) (Controller, error)
	Save() error
	Load() error
}

type Controller interface {
	Views() (vs []string)
	Upload(w io.Writer, files []string) (err error)
	List(w io.Writer, view string) (err error)
	Edit(w io.Writer, view, key string) (err error)
	Apply(view, key string, vals map[string][]string) ([]*forms.Updated, error)
	Query(view, key string) (err error)
}
