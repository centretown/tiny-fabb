package monitor

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Operator int

const (
	Open Operator = iota
	Close
)

var (
	ErrTimeout error = errors.New("timed out")
)

// Bus - i/o communications
type Bus struct {
	In        chan string
	Out       chan string
	Operation chan Operator
	Done      chan bool
	Timeout   time.Duration
	Err       chan error
}

// NewBus creates an i/o bus
func NewBus() (bus *Bus) {
	bus = &Bus{
		In:      make(chan string, 12),
		Out:     make(chan string),
		Done:    make(chan bool),
		Err:     make(chan error),
		Timeout: time.Millisecond * 500,
	}
	return
}

func (bus *Bus) Capture(out string, terminators ...string) (results []string, err error) {

	results = make([]string, 0)
	bus.Out <- out
	limit := time.Now().Add(bus.Timeout)

	for {
		select {
		case in := <-bus.In:
			limit = limit.Add(bus.Timeout)
			in = strings.TrimSpace(in)
			results = append(results, in)
			if Terminated(in, terminators) {
				return
			}
		case err = <-bus.Err:
			return
		default:
			if time.Now().After(limit) {
				err = fmt.Errorf("timeout processing %s", out)
				return
			}
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
