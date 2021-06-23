package grbl

import (
	"fmt"
	"html/template"
	"os"
	"sort"
	"testing"

	"github.com/centretown/tiny-fabb/web"
)

func testWeb(t *testing.T) {
	layout := template.Must(template.ParseFiles("../assets/layout.html"))
	gctl := NewController(layout)
	gctl.Settings.StepPortInvertMask = 6
	gctl.Settings.DirPortInvertMask = 5
	gctl.Settings.StatusReportMask = 2
	gctl.Settings.HomingDirInvert = 6
	gctl.Settings.StepPulse = 123
	gctl.Settings.LaserMode = true
	gctl.Settings.StepEnableInvert = true
	gctl.Settings.LimitPinsInvert = true
	gctl.Settings.ProbePinInvert = true
	err := gctl.List(os.Stdout, "settings")
	if err != nil {
		t.Fatal(err)
	}

	settings := gctl.views["settings"]
	var ids web.WebIds = make(web.WebIds, 0, len(settings))
	for id, _ := range settings { //NOLINT
		ids = append(ids, id)
	}

	sort.Sort(ids)

	for _, id := range ids {
		gv := settings[id]
		first := gv.Entries[0]
		for _, ge := range gv.Entries {
			f := ge.FormatInput(gv.Value, first)
			fmt.Println(f)
		}
	}

	for _, id := range ids {
		err = gctl.Edit(os.Stdout, "settings", web.WebId(id).String())
		if err != nil {
			t.Fatal(err)
		}
	}

	var vals = map[string][]string{}
	// gctl.Settings.StepPulse = 123
	vals[idStepPulse.String()] = []string{"888"}
	gctl.Apply("settings", idStepPulse.String(), vals)
	if gctl.Settings.StepPulse != 888 {
		t.Fatal(gctl.Settings.StepPulse)
	}
	t.Log(gctl.Settings.StepPulse)

	// gctl.Settings.StepPortInvertMask = 6
	vals[idStepPortInvertMask.Index(1)] = []string{"true"}
	vals[idStepPortInvertMask.Index(2)] = []string{"false"}
	vals[idStepPortInvertMask.Index(3)] = []string{"true"}
	gctl.Apply("settings", idStepPortInvertMask.String(), vals)
	if gctl.Settings.StepPortInvertMask != 5 {
		t.Fatal(gctl.Settings.StepPortInvertMask)
	}
	t.Log(gctl.Settings.StepPortInvertMask)

	// gctl.Settings.DirPortInvertMask = 5
	vals[idDirPortInvertMask.Index(1)] = []string{"false"}
	vals[idDirPortInvertMask.Index(2)] = []string{"true"}
	vals[idDirPortInvertMask.Index(3)] = []string{"false"}
	gctl.Apply("settings", idDirPortInvertMask.String(), vals)
	if gctl.Settings.DirPortInvertMask != 2 {
		t.Fatal(gctl.Settings.DirPortInvertMask)
	}
	t.Log(gctl.Settings.DirPortInvertMask)

	vals[idStatusReportMask.String()] = []string{"0"}
	vals[idStatusReportMask.Index(3)] = []string{"true"}
	gctl.Apply("settings", idStatusReportMask.String(), vals)
	if gctl.Settings.StatusReportMask != 2 {
		t.Fatal(gctl.Settings.StatusReportMask)
	}
	t.Log(gctl.Settings.StatusReportMask)

	vals[idStatusReportMask.String()] = []string{"1"}
	vals[idStatusReportMask.Index(3)] = []string{"false"}
	gctl.Apply("settings", idStatusReportMask.String(), vals)
	if gctl.Settings.StatusReportMask != 1 {
		t.Fatal(gctl.Settings.StatusReportMask)
	}
	t.Log(gctl.Settings.StatusReportMask)

	vals[idStatusReportMask.String()] = []string{"1"}
	vals[idStatusReportMask.Index(3)] = []string{"true"}
	gctl.Apply("settings", idStatusReportMask.String(), vals)
	if gctl.Settings.StatusReportMask != 3 {
		t.Fatal(gctl.Settings.StatusReportMask)
	}
	t.Log(gctl.Settings.StatusReportMask)

	vals[idHomingDirInvert.Index(1)] = []string{"true"}
	vals[idHomingDirInvert.Index(2)] = []string{"false"}
	vals[idHomingDirInvert.Index(3)] = []string{"false"}
	vals[idHomingDirInvert.Index(6)] = []string{"true"}
	gctl.Apply("settings", idHomingDirInvert.String(), vals)
	if gctl.Settings.HomingDirInvert != 33 {
		t.Fatal(gctl.Settings.HomingDirInvert)
	}
	t.Log(gctl.Settings.HomingDirInvert)

	// gctl.Settings.LaserMode = true
	vals[idLaserMode.String()] = []string{"false"}
	gctl.Apply("settings", idLaserMode.String(), vals)
	if gctl.Settings.LaserMode != false {
		t.Fatal(gctl.Settings.LaserMode)
	}
	t.Log(gctl.Settings.LaserMode)

	vals[idLaserMode.String()] = []string{"true"}
	gctl.Apply("settings", idLaserMode.String(), vals)
	if gctl.Settings.LaserMode != true {
		t.Fatal(gctl.Settings.LaserMode)
	}
	t.Log(gctl.Settings.LaserMode)

	// gctl.Settings.StepEnableInvert = true
	// gctl.Settings.LimitPinsInvert = true
	// gctl.Settings.ProbePinInvert = true

}
