package monitor

import (
	"strings"
	"time"
)

// Bus - i/o communications
type Bus struct {
	In   chan string
	Out  chan string
	Done chan bool
	Err  chan error
}

// NewBus creates an i/o bus
func NewBus() (bus *Bus) {
	bus = &Bus{
		In:   make(chan string, 12),
		Out:  make(chan string),
		Done: make(chan bool),
		Err:  make(chan error),
	}
	return
}

func (bus *Bus) Capture(out string, terminators ...string) (results []string, err error) {
	results = make([]string, 0)
	bus.Out <- out

	for {
		select {
		case in := <-bus.In:
			in = strings.TrimSpace(in)
			results = append(results, in)
			if Terminated(in, terminators) {
				return
			}
		case err = <-bus.Err:
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

func Terminated(s string, terminators []string) bool {
	if len(terminators) < 1 {
		return true
	}
	for _, terminator := range terminators {
		if strings.HasPrefix(s, terminator) {
			return true
		}
	}
	return false
}
