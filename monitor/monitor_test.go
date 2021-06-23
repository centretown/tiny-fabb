package monitor

import (
	"testing"

	"github.com/centretown/tiny-fabb/serialio"
)

func TestMonitor(t *testing.T) {
	// serialio.EnumerateSerialPorts()
	ports := serialio.ListSerial()
	t.Log(ports)
	for _, port := range ports {
		t.Log()
		t.Log(port)
		sio, err := serialio.GetSerialIO(port)
		if err != nil {
			t.Log(err)
			continue
		}

		bus := NewBus()
		go Monitor(sio, bus)
		capture(t, bus)
	}
}

func capture(t *testing.T, bus *Bus) {
	displayResults := func(results []string, err error) error {
		for _, result := range results {
			t.Log(result)
		}
		if err != nil {
			t.Fatal(err)
		}
		return err
	}
	terminators := []string{"ok", "error"}
	displayResults(bus.Capture("$$", terminators...))
	displayResults(bus.Capture("$131", terminators...))
	displayResults(bus.Capture("$I", terminators...))
	displayResults(bus.Capture("$#", terminators...))
	displayResults(bus.Capture("?", terminators...))
	displayResults(bus.Capture("$C", terminators...))
	displayResults(bus.Capture("$C", terminators...))
	displayResults(bus.Capture("$", terminators...))
	displayResults(bus.Capture("", terminators...))
}
