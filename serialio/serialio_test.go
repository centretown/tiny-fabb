package serialio

import (
	"testing"
	"time"

	"github.com/centretown/tiny-fabb/monitor"
)

func TestSerialIo(t *testing.T) {
	prv := &Provider{}
	testList(t, prv, 2)
	testList(t, prv, 2, "ACM")
	testList(t, prv, 2, "ACM", "USB")
}

func testList(t *testing.T, prv *Provider, count int, filter ...string) {
	t.Log(filter)

	prv.Update(filter...)
	for i := 0; i < count; i++ {
		t.Logf("available %v\n", prv.List())
		time.Sleep(time.Second)
	}
}

func TestMonitor(t *testing.T) {
	prv := &Provider{}
	testProvider(t, prv)
}

func testProvider(t *testing.T, prv monitor.Provider) {

	prv.Update()
	ports := prv.List()
	t.Log(ports)
	for _, port := range ports {
		t.Log()
		t.Log(port)
		mio, err := prv.Get(port)
		if err != nil {
			t.Log(err)
			continue
		}
		bus := monitor.NewBus()
		go monitor.Monitor(mio, bus)
		capture(t, bus)

	}
}

func capture(t *testing.T, bus *monitor.Bus) {
	displayResults := func(results []string, err error) {
		if err != nil {
			t.Log(err)
		} else {
			for _, result := range results {
				t.Log(result)
			}
		}
	}
	terminators := []string{"ok", "error"}
	displayResults(bus.Capture("$I", terminators...))
	displayResults(bus.Capture("$#", terminators...))
	displayResults(bus.Capture("?", terminators...))
	displayResults(bus.Capture("$C", terminators...))
	displayResults(bus.Capture("$C", terminators...))
	displayResults(bus.Capture("$", terminators...))
	displayResults(bus.Capture("", terminators...))
	displayResults(bus.Capture("$$", terminators...))
	displayResults(bus.Capture("$0=10", terminators...))
}
