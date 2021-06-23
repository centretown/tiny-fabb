package grbl

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"testing"

	"github.com/centretown/tiny-fabb/serialio"
	"github.com/centretown/tiny-fabb/web"
)

func testMonitor(t *testing.T) {
	ports := serialio.ListSerial()
	t.Log(ports)
	layout := template.Must(template.ParseFiles("../assets/layout.html"))
	gctl := NewController(layout)
	gctl.Title = "GRBL ESP32a"
	gctl.Port = "/dev/ttyUSB0"

	displayResults := func(results []string, err error) error {
		for _, result := range results {
			var (
				key web.WebId
				val string
			)

			if strings.HasPrefix(result, "ok") {
				break
			}
			if strings.HasPrefix(result, "error") {
				t.Fatal(result)
				break
			}

			fmt.Sscanf(result, "$%d=%s", &key, &val)
			t.Log(key, val)
			form, err := gctl.getForm("settings", key.String())
			if err != nil {
				t.Log(err)
				continue
			}
			ent := form.Entries[0]
			err = ent.ScanInput(val, form.Value)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(val, ent.Value(form.Value))
		}
		if err != nil {
			t.Fatal(err)
		}
		return err
	}

	err := gctl.ActivateController()
	if err != nil {
		t.Fatal(err)
	}

	displayResults(gctl.bus.Capture("$$", "ok", "error"))
}

func TestQuery(t *testing.T) {
	ports := serialio.ListSerial()
	t.Log(ports)

	layout := template.Must(template.ParseFiles("../assets/layout.test"))
	gctl := NewController(layout)
	gctl.Title = "GRBL ESP32a"
	gctl.Port = "/dev/ttyUSB1"
	err := gctl.ActivateController()
	if err != nil {
		t.Fatal(err)
	}

	err = gctl.List(os.Stdout, "settings")
	if err != nil {
		t.Fatal(err)
	}

	err = gctl.Query("settings", idSettings.String())
	if err != nil {
		t.Fatal(err)
	}

	err = gctl.List(os.Stdout, "settings")
	if err != nil {
		t.Fatal(err)
	}

	err = gctl.Query("commands", idParameters.String())
	if err != nil {
		t.Fatal(err)
	}

	err = gctl.Query("commands", idParserState.String())
	if err != nil {
		t.Fatal(err)
	}

	err = gctl.Query("commands", idBuildInfo.String())
	if err != nil {
		t.Fatal(err)
	}

	err = gctl.Query("commands", idStartupBlocks.String())
	if err != nil {
		t.Fatal(err)
	}

	err = gctl.Query("commands", idCodeMode.String())
	if err != nil {
		t.Fatal(err)
	}

	err = gctl.Query("commands", idCodeMode.String())
	if err != nil {
		t.Fatal(err)
	}

	err = gctl.List(os.Stdout, "commands")
	if err != nil {
		t.Fatal(err)
	}
}
