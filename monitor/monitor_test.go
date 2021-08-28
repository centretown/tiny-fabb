package monitor

// func testProvider(t *testing.T, prv Provider) {
// 	prv.Update()
// 	ports := prv.List()
// 	t.Log(ports)
// 	for _, port := range ports {
// 		t.Log()
// 		t.Log(port)
// 		sio, err := prv.Get(port)
// 		if err != nil {
// 			t.Log(err)
// 			continue
// 		}

// 		io, ok := sio.(MonitorIO)
// 		if ok {
// 			bus := NewBus()
// 			go Monitor(io, bus)
// 			capture(t, bus)

// 		}

// 	}
// }

// func TestMonitor(t *testing.T) {
// 	// serialio.EnumerateSerialPorts()
// 	prv := serialio.NewProvider()
// 	prv.Setup()
// 	testProvider(t, prv)
// }

// func capture(t *testing.T, bus *Bus) {
// 	displayResults := func(results []string, err error) error {
// 		for _, result := range results {
// 			t.Log(result)
// 		}
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		return err
// 	}
// 	terminators := []string{"ok", "error"}
// 	displayResults(bus.Capture("$131", terminators...))
// 	displayResults(bus.Capture("$I", terminators...))
// 	displayResults(bus.Capture("$#", terminators...))
// 	displayResults(bus.Capture("?", terminators...))
// 	displayResults(bus.Capture("$C", terminators...))
// 	displayResults(bus.Capture("$C", terminators...))
// 	displayResults(bus.Capture("$", terminators...))
// 	displayResults(bus.Capture("", terminators...))
// 	displayResults(bus.Capture("$$", terminators...))

// 	// fmt.Println("5 seconds")
// 	// time.Sleep(time.Second * 5)

// 	// // bus.Operation <- Close
// 	// displayResults(bus.Capture("$$", terminators...))

// }
