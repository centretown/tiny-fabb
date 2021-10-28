package camera

import "github.com/centretown/tiny-fabb/forms"

type Streamer interface {
	// Get(id string) (string, error)
	// Set(id, val string) error
	BindProperties() forms.Forms
	UpdateProperties() error
	SetProperty(ent *forms.Entry, val string) error

	Open() error
	Read() ([]byte, error)
	Close() error
}
