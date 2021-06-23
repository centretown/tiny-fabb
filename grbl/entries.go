package grbl

import (
	"fmt"

	"github.com/centretown/tiny-fabb/web"
)

var commandEntries = map[web.WebId]web.Entries{
	idSettings: {
		{
			ID:    idSettings.String(),
			Code:  "$$",
			Label: "View and write Grbl settings",
			URL:   "#and-xval---view-and-write-grbl-settings",
			Type:  "text"}},
	// Scan:  scanSetting}},
	idParameters: {
		{
			ID:    idParameters.String(),
			Code:  "$#",
			Label: "View gcode parameters",
			URL:   "#---view-gcode-parameters",
			Type:  "text"}},
	idParserState: {
		{
			ID:    idParserState.String(),
			Code:  "$G",
			Label: "View gcode parser state",
			URL:   "#g---view-gcode-parser-state",
			Type:  "text"}},
	idBuildInfo: {
		{
			ID:    idBuildInfo.String(),
			Code:  "$I",
			Label: "View build info",
			URL:   "#i---view-build-info",
			Type:  "text"}},
	// Scan:  scanInfo}},
	idStartupBlocks: {
		{
			ID:    idStartupBlocks.String(),
			Code:  "$N",
			Label: "View startup blocks",
			URL:   "#n---view-startup-blocks",
			Type:  "text"}},
	idCodeMode: {
		{
			ID:    idCodeMode.String(),
			Code:  "$C",
			Label: "Check gcode mode",
			URL:   "#c---check-gcode-mode",
			Type:  "text"}},
	idKillAlarm: {
		{
			ID:    idKillAlarm.String(),
			Code:  "$X",
			Label: "Kill alarm lock",
			URL:   "#x---kill-alarm-lock",
			Type:  "text"}},
	idRunHomingCycle: {
		{
			ID:    idRunHomingCycle.String(),
			Code:  "$H",
			Label: "Run homing cycle",
			URL:   "#h---run-homing-cycle",
			Type:  "text"}},
	idRunJoggingCycle: {
		{
			ID:    idRunJoggingCycle.String(),
			Code:  "$J",
			Label: "Run jogging motion",
			URL:   "#jline---run-jogging-motion",
			Type:  "text"}},
	idEraseRestore: {
		{
			ID:    idEraseRestore.String(),
			Code:  "$RST=$",
			Label: "Erase and restore",
			URL:   "#rst-rst-and-rst--restore-grbl-settings-and-data-to-defaults",
			Type:  "text"}},
	idEraseZero: {
		{
			ID:    idEraseZero.String(),
			Code:  "$RST=#",
			Label: "Erase and zero",
			URL:   "#rst-rst-and-rst--restore-grbl-settings-and-data-to-defaults",
			Type:  "text"}},
	idClearRestore: {
		{
			ID:    idClearRestore.String(),
			Code:  "$RST=*",
			Label: "Clear and restore",
			URL:   "#rst-rst-and-rst--restore-grbl-settings-and-data-to-defaults",
			Type:  "text"}},
}

const (
	microSec       = "(\xc2\xb5s)"
	axisEntryCount = 7
)

var (
	axis = []string{"X", "Y", "Z", "A", "B", "C"}
)

func axisMasks(id web.WebId, code string, label string, url string) (entries []*web.Entry) {
	entries = make([]*web.Entry, axisEntryCount)
	entries[0] = &web.Entry{
		ID:    id.String(),
		Code:  code,
		Label: label,
		URL:   url,
		Type:  "mask",
	}

	for i := 1; i < axisEntryCount; i++ {
		entries[i] = axisMask(id, i, url)
	}
	return
}

func axisMask(id web.WebId, index int, url string) (ent *web.Entry) {
	var mask uint = 1 << (index - 1)
	ent = &web.Entry{
		ID:    id.Index(index),
		Label: "Invert " + axis[index-1],
		Type:  "checkbox",
		URL:   url,
		Val: func(v interface{}) interface{} {
			return web.Mask(v, mask)
		},
		Scan: func(s string, v interface{}) (err error) {
			err = web.UnMasks(v, mask, s)
			return
		},
	}
	return
}

// setting formatting hints
var settingEntries = map[web.WebId]web.Entries{
	idStepPulse: {
		{
			ID:    idStepPulse.String(),
			Code:  "$0",
			Label: "Step Pulse " + microSec,
			URL:   "#0--step-pulse-microseconds",
			Type:  "number",
			Min:   3,
			Max:   1000,
		}},
	idStepIdleDelay: {
		{
			ID:    idStepIdleDelay.String(),
			Code:  "$1",
			Label: "Step Idle Delay (ms)",
			URL:   "#1---step-idle-delay-milliseconds",
			Type:  "number",
			Min:   0,
			Max:   1000,
		}},
	idStepPortInvertMask: axisMasks(idStepPortInvertMask,
		"$2", "Invert Step Port", "#2--step-port-invert-mask"),
	idDirPortInvertMask: axisMasks(idDirPortInvertMask,
		"$3", "Invert Direction Port", "#3--direction-port-invert-mask"),
	idStepEnableInvert: {
		{
			ID:    idStepEnableInvert.String(),
			Code:  "$4",
			Label: "Invert Step Enable",
			URL:   "#4---step-enable-invert-boolean",
			Type:  "checkbox"}},
	idLimitPinsInvert: {
		{
			ID:    idLimitPinsInvert.String(),
			Code:  "$5",
			Label: "Invert Limit Pins",
			URL:   "#5----limit-pins-invert-boolean",
			Type:  "checkbox"}},
	idProbePinInvert: {
		{
			ID:    idProbePinInvert.String(),
			Code:  "$6",
			Label: "Invert Probe Pin",
			URL:   "#6----probe-pin-invert-boolean",
			Type:  "checkbox"}},

	idStatusReportMask: {
		{
			ID:    idStatusReportMask.String(),
			Code:  "$10",
			Label: "Status Report",
			URL:   "#10---status-report-mask",
			Type:  "mask",
			Scan: func(s string, v interface{}) (err error) {
				var choice int
				fmt.Sscan(s, &choice)
				err = web.UnMask(v, 1, choice == 1)
				return
			},
		},
		{
			ID:    idStatusReportMask.Index(1),
			Label: "Show Work Position",
			Type:  "radio",
			URL:   "#10---status-report-mask",
			Item:  0,
			Val: func(v interface{}) (r interface{}) {
				r = !web.Mask(v, 1)
				return
			},
		},
		{
			ID:    idStatusReportMask.Index(2),
			Label: "Show Machine Position",
			Type:  "radio",
			URL:   "#10---status-report-mask",
			Item:  1,
			Val: func(v interface{}) (r interface{}) {
				r = web.Mask(v, 1)
				return
			},
		},
		{
			ID:    idStatusReportMask.Index(3),
			Label: "Data planner, Serial RX available data",
			Type:  "checkbox",
			URL:   "#10---status-report-mask",
			Val: func(v interface{}) (r interface{}) {
				r = web.Mask(v, 2)
				return
			},
			Scan: func(s string, v interface{}) (err error) {
				err = web.UnMasks(v, 2, s)
				return
			},
		},
	},
	idJunctionDeviation: {
		{
			ID:    idJunctionDeviation.String(),
			Code:  "$11",
			Label: "Junction Deviation (mm)",
			URL:   "#11---junction-deviation-mm",
			Type:  "number",
			Min:   0,
			Max:   1.00,
			Step:  0.01,
		},
	},
	idArcTolerance: {
		{
			ID:    idArcTolerance.String(),
			Code:  "$12",
			Label: "Arc Tolerance (mm)",
			URL:   "#12--arc-tolerance-mm",
			Type:  "number",
			Min:   0,
			Max:   1.000,
			Step:  0.001,
		},
	},
	idReportInches: {
		{
			ID:    idReportInches.String(),
			Code:  "$13",
			Label: "Report Using Inches",
			URL:   "#13---report-inches-boolean",
			Type:  "checkbox"},
	},

	idSoftLimits: {
		{
			ID:    idSoftLimits.String(),
			Code:  "$20",
			Label: "Enable Soft Limits",
			URL:   "#20---soft-limits-boolean",
			Type:  "checkbox"},
	},
	idHardLimits: {
		{
			ID:    idHardLimits.String(),
			Code:  "$21",
			Label: "Enable Hard Limits",
			URL:   "#21---hard-limits-boolean",
			Type:  "checkbox",
		}},
	idHomingCycle: {
		{
			ID:    idHomingCycle.String(),
			Code:  "$22",
			Label: "Enable Homing Cycle",
			URL:   "#22---homing-c1ycle-boolean",
			Type:  "checkbox",
		}},
	idHomingDirInvert: axisMasks(idHomingDirInvert,
		"$23", "Invert Homing Direction", "#23---homing-dir-invert-mask"),
	idHomingFeed: {
		{
			ID:    idHomingFeed.String(),
			Code:  "$24",
			Label: "Homing feed (mm/min)",
			URL:   "#24---homing-feed-mmmin",
			Type:  "number",
		}},
	idHomingSeek: {
		{
			ID:    idHomingSeek.String(),
			Code:  "$25",
			Label: "Homing seek (mm/min)",
			URL:   "#25---homing-seek-mmmin",
			Type:  "number"}},
	idHomingDebounce: {
		{
			ID:    idHomingDebounce.String(),
			Code:  "$26",
			Label: "Homing debounce (ms)",
			URL:   "#26---homing-debounce-milliseconds",
			Type:  "number"}},
	idHomingPulloff: {
		{
			ID:    idHomingPulloff.String(),
			Code:  "$27",
			Label: "Homing pull-off (mm)",
			URL:   "#27---homing-pull-off-mm",
			Type:  "number"}},

	idMaxSpindleSpeed: {
		{
			ID:    idMaxSpindleSpeed.String(),
			Code:  "$30",
			Label: "Max spindle speed (rpm)",
			URL:   "#30---max-spindle-speed-rpm",
			Type:  "number"}},
	idMinSpindleSpeed: {
		{
			ID:    idMinSpindleSpeed.String(),
			Code:  "$31",
			Label: "Min spindle speed (rpm)",
			URL:   "#31---min-spindle-speed-rpm",
			Type:  "number"}},
	idLaserMode: {
		{
			ID:    idLaserMode.String(),
			Code:  "$32",
			Label: "Laser mode, boolean",
			URL:   "#32---laser-mode-boolean",
			Type:  "checkbox"}},

	idStepsX: {
		{
			ID:    idStepsX.String(),
			Code:  "$100",
			Label: "X (steps/mm)",
			URL:   "#100-101-and-102--xyz-stepsmm",
			Type:  "number"}},
	idStepsY: {
		{
			ID:    idStepsY.String(),
			Code:  "$101",
			Label: "Y (steps/mm)",
			URL:   "#100-101-and-102--xyz-stepsmm",
			Type:  "number"}},
	idStepsZ: {
		{
			ID:    idStepsZ.String(),
			Code:  "$102",
			Label: "Z (steps/mm)",
			URL:   "#100-101-and-102--xyz-stepsmm",
			Type:  "number"}},

	idStepsA: {
		{
			ID:    idStepsA.String(),
			Code:  "$103",
			Label: "A (steps/mm)",
			URL:   "#100-101-and-102--xyz-stepsmm",
			Type:  "number"}},
	idStepsB: {
		{
			ID:    idStepsB.String(),
			Code:  "$104",
			Label: "B (steps/mm)",
			URL:   "#100-101-and-102--xyz-stepsmm",
			Type:  "number"}},
	idStepsC: {
		{
			ID:    idStepsC.String(),
			Code:  "$105",
			Label: "C (steps/mm)",
			URL:   "#100-101-and-102--xyz-stepsmm",
			Type:  "number"}},

	idMaxRateX: {
		{
			ID:    idMaxRateX.String(),
			Code:  "$110",
			Label: "X Max rate (mm/min)",
			URL:   "#110-111-and-112--xyz-max-rate-mmmin",
			Type:  "number"}},
	idMaxRateY: {
		{
			ID:    idMaxRateY.String(),
			Code:  "$111",
			Label: "Y Max rate (mm/min)",
			URL:   "#110-111-and-112--xyz-max-rate-mmmin",
			Type:  "number"}},
	idMaxRateZ: {
		{
			ID:    idMaxRateZ.String(),
			Code:  "$112",
			Label: "Z Max rate (mm/min)",
			URL:   "#110-111-and-112--xyz-max-rate-mmmin",
			Type:  "number"}},

	idMaxRateA: {
		{
			ID:    idMaxRateA.String(),
			Code:  "$113",
			Label: "A Max rate (mm/min)",
			URL:   "#110-111-and-112--xyz-max-rate-mmmin",
			Type:  "number"}},
	idMaxRateB: {
		{
			ID:    idMaxRateB.String(),
			Code:  "$114",
			Label: "B Max rate (mm/min)",
			URL:   "#110-111-and-112--xyz-max-rate-mmmin",
			Type:  "number"}},
	idMaxRateC: {
		{
			ID:    idMaxRateC.String(),
			Code:  "$115",
			Label: "C Max rate (mm/min)",
			URL:   "#110-111-and-112--xyz-max-rate-mmmin",
			Type:  "number"}},

	idAccelX: {
		{
			ID:    idAccelX.String(),
			Code:  "$120",
			Label: "X Acceleration (mm/sec\xc2\xb2)",
			URL:   "#120-121-122--xyz-acceleration-mmsec2",
			Type:  "number"}},
	idAccelY: {
		{
			ID:    idAccelY.String(),
			Code:  "$121",
			Label: "Y Acceleration (mm/sec\xc2\xb2)",
			URL:   "#120-121-122--xyz-acceleration-mmsec2",
			Type:  "number"}},
	idAccelZ: {
		{
			ID:    idAccelZ.String(),
			Code:  "$122",
			Label: "Z Acceleration (mm/sec\xc2\xb2)",
			URL:   "#120-121-122--xyz-acceleration-mmsec2",
			Type:  "number"}},
	idAccelA: {
		{
			ID:    idAccelA.String(),
			Code:  "$123",
			Label: "A Acceleration (mm/sec\xc2\xb2)",
			URL:   "#120-121-122--xyz-acceleration-mmsec2",
			Type:  "number"}},
	idAccelB: {
		{
			ID:    idAccelB.String(),
			Code:  "$124",
			Label: "B Acceleration (mm/sec\xc2\xb2)",
			URL:   "#120-121-122--xyz-acceleration-mmsec2",
			Type:  "number"}},
	idAccelC: {
		{
			ID:    idAccelC.String(),
			Code:  "$125",
			Label: "C Acceleration (mm/sec\xc2\xb2)",
			URL:   "#120-121-122--xyz-acceleration-mmsec2",
			Type:  "number"}},
	idMaxTravelX: {
		{
			ID:    idMaxTravelX.String(),
			Code:  "$130",
			Label: "X Max travel (mm)",
			URL:   "#130-131-132--xyz-max-travel-mm",
			Type:  "number"}},
	idMaxTravelY: {
		{
			ID:    idMaxTravelY.String(),
			Code:  "$131",
			Label: "Y Max travel (mm)",
			URL:   "#130-131-132--xyz-max-travel-mm",
			Type:  "number"}},
	idMaxTravelZ: {
		{
			ID:    idMaxTravelZ.String(),
			Code:  "$132",
			Label: "Z Max travel (mm)",
			URL:   "#130-131-132--xyz-max-travel-mm",
			Type:  "number"}},
	idMaxTravelA: {
		{
			ID:    idMaxTravelX.String(),
			Code:  "$133",
			Label: "A Max travel (mm)",
			URL:   "#130-131-132--xyz-max-travel-mm",
			Type:  "number"}},
	idMaxTravelB: {
		{
			ID:    idMaxTravelB.String(),
			Code:  "$134",
			Label: "B Max travel (mm)",
			URL:   "#130-131-132--xyz-max-travel-mm",
			Type:  "number"}},
	idMaxTravelC: {
		{
			ID:    idMaxTravelC.String(),
			Code:  "$135",
			Label: "C Max travel (mm)",
			URL:   "#130-131-132--xyz-max-travel-mm",
			Type:  "number"}},
}
