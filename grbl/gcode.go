package grbl

import (
	"github.com/centretown/tiny-fabb/forms"
)

const (
	Parameters forms.WebId = iota
	Motion
	PlaneSelection
	Diameter
	DistanceMode
	FeedRateMode
	Units
	CutterRadiusCompensation
	ToolLengthOffset
	ReturnModeInCannedCycles
	CoordinateSystemSelection
	Stopping
	ToolChange
	SpindleTurning
	Coolant
	OverrideSwitches
	FlowControl
	NonModal
)

type Group struct {
	ID    forms.WebId
	Label string
}

var Groups = map[forms.WebId]*forms.Entry{
	Motion: {
		ID:    Motion.String(),
		Label: "Motion ('Group 1')",
	},
	PlaneSelection: {
		ID:    PlaneSelection.String(),
		Label: "Plane selection",
	},
	Diameter: {
		ID:    Diameter.String(),
		Label: "Diameter / Radius for lathes",
	},
	DistanceMode: {
		ID:    DistanceMode.String(),
		Label: "Distance Mode",
	},
	FeedRateMode: {
		ID:    FeedRateMode.String(),
		Label: "Feed Rate Mode",
	},
	Units: {
		ID:    Units.String(),
		Label: "Units",
	},
	CutterRadiusCompensation: {
		ID:    CutterRadiusCompensation.String(),
		Label: "Cutter Radius Compensation",
	},
	ToolLengthOffset: {
		ID:    ToolLengthOffset.String(),
		Label: "Tool Length Offset",
	},
	ReturnModeInCannedCycles: {
		ID:    ReturnModeInCannedCycles.String(),
		Label: "Return Mode in Canned Cycles",
	},
	CoordinateSystemSelection: {
		ID:    CoordinateSystemSelection.String(),
		Label: "Coordinate System Selection",
	},
	Stopping: {
		ID:    Stopping.String(),
		Label: "Stopping",
	},
	ToolChange: {
		ID:    ToolChange.String(),
		Label: "Tool Change",
	},
	SpindleTurning: {
		ID:    SpindleTurning.String(),
		Label: "Spindle Turning",
	},
	Coolant: {
		ID:    Coolant.String(),
		Label: "Coolant",
	},
	OverrideSwitches: {
		ID:    OverrideSwitches.String(),
		Label: "Override Switches",
	},
	FlowControl: {
		ID:    FlowControl.String(),
		Label: "Flow Control",
	},
	NonModal: {
		ID:    NonModal.String(),
		Label: "Non-modal codes ('Group 0')",
	},
}

type GCode struct {
	Group      *forms.Entry
	Code       string
	Label      string
	Notes      string
	Parameters []*GCode
}

var GCodes = map[string]*GCode{
	"G0": {
		Group: Groups[Motion],
		Code:  "G0",
		Label: "Rapid positioning",
		Notes: "Switch to rapid linear motion mode (seek). Used to get the tool somewhere quickly without cutting --- moves the machine as quickly as possible along each axis --- an axis which needs less movement will finish before the others, so one cannot count on the movement being a straight line.",
	},
	"G1": {
		Group: Groups[Motion],
		Code:  "G1",
		Label: "Linear interpolation",
		Notes: "Switch to linear motion at the current feed rate. Used to cut a straight line --- the interpreter will determine the acceleration needed along each axis to ensure direct movement from the original to the destination point at no more than the current Feed rate (F see below).",
	},
	"G2": {
		Group: Groups[Motion],
		Code:  "G2",
		Label: "Circular interpolation, clockwise",
		Notes: "Switch to clockwise arc mode. The interpreter will cut an arc or circle from the current position to the destination using the specified radius (R) or center (IJK location) at the current Feed rate (F see below) in the plane selected by G17/18/19.",
	},
	"G3": {
		Group: Groups[Motion],
		Code:  "G3",
		Label: "Circular interpolation, counterclockwise",
		Notes: "Switch to anti-clockwise arc mode. Corollary to G2 above.",
	},
	"G4": {
		Group: Groups[Motion],
		Code:  "G4",
		Label: "Dwell",
		Notes: "This should probably be calculated to be only one or two spindle rotations for best efficiency. Dwell time is expressed using a parameter (may be X, U, or P) which determines the time unit (seconds, milliseconds, &c.) P, for seconds, is supported and used by Grbl, typically X and U express the duration in milliseconds.",
	},
	"G38.2": {
		Group: Groups[Motion],
		Code:  "G38.2",
		Label: "Straight Probe",
		Notes: "Probe toward workpiece, stop on contact, signal error if failure.",
	},
	"G38.3": {
		Group: Groups[Motion],
		Code:  "G38.3",
		Label: "Probe",
		Notes: "Probe toward workpiece, stop on contact.",
	},
	"G38.4": {
		Group: Groups[Motion],
		Code:  "G38.4",
		Label: "Probe",
		Notes: "Probe away workpiece, stop on contact, signal error if failure.",
	},
	"G38.5": {
		Group: Groups[Motion],
		Code:  "G38.5",
		Label: "Probe",
		Notes: "Probe away workpiece, stop on contact.",
	},
	"G80": {
		Group: Groups[Motion],
		Code:  "G80",
		Label: "Motion mode cancel",
		Notes: "Canned cycle",
	},
	"G54": {
		Group: Groups[CoordinateSystemSelection],
		Code:  "G54",
		Label: "Fixture offset 1",
		Notes: "Fixture offset 1--6. CF G10 and G92.[21] Note that G54 is reserved by Carbide Motion, and will be reset by the software.[22]",
	},
	"G55": {
		Group: Groups[CoordinateSystemSelection],
		Code:  "G55",
		Label: "Fixture offset 2",
	},
	"G56": {
		Group: Groups[CoordinateSystemSelection],
		Code:  "G56",
		Label: "Fixture offset 3",
	},
	"G57": {
		Group: Groups[CoordinateSystemSelection],
		Code:  "G57",
		Label: "Fixture offset 4",
	},
	"G58": {
		Group: Groups[CoordinateSystemSelection],
		Code:  "G58",
		Label: "Fixture offset 5",
	},
	"G59": {
		Group: Groups[CoordinateSystemSelection],
		Code:  "G59",
		Label: "Fixture offset 6",
	},
	"G17": {
		Group: Groups[PlaneSelection],
		Code:  "G17",
		Label: "Select the XY plane (for arcs)",
		Notes: "Use I and J",
	},
	"G18": {
		Group: Groups[PlaneSelection],
		Code:  "G18",
		Label: "Select the XZ plane (for arcs)",
		Notes: "Use I and K",
	},
	"G19": {
		Group: Groups[PlaneSelection],
		Code:  "G19",
		Label: "Select the YZ plane (for arcs)",
		Notes: "Use J and K",
	},
	"G90": {
		Group: Groups[DistanceMode],
		Code:  "G90",
		Label: "Switch to absolute distance mode",
		Notes: "Coordinates are now relative to the origin of the currently active coordinate system, as opposed to the current position. G0 X-10 Y5 will move to the position 10 units to the left and 5 above the origin X0,Y0. cf. G91 below.",
	},
	"G91": {
		Group: Groups[DistanceMode],
		Code:  "G91",
		Label: "Switch to incremental distance mode",
		Notes: "Coordinates are now relative to the current position, with no consideration for machine origin. G0 X-10 Y5 will move to the position 10 units to the left and 5 above the current position. cf. G90 above.",
	},
	"G93": {
		Group: Groups[FeedRateMode],
		Code:  "G93",
		Label: "Set inverse time feed rate mode",
		Notes: "An F word is interpreted to mean that the move should be completed in (one divided by the F number) minutes. For example, if F is 2, the move should be completed in half a minute.",
	},
	"G94": {
		Group: Groups[FeedRateMode],
		Code:  "G94",
		Label: "Set units per minute feed rate mode",
		Notes: "An F Word is interpreted to mean the controlled point should move at a certain number of units (or degrees) per minute.",
	},
	"G20": {
		Group: Groups[Units],
		Code:  "G20",
		Label: "Units will be in inches",
		Notes: "Best practice: do this at the start of a program and nowhere else. The usual minimum increment in G20 is one ten-thousandth of an inch (0.0001\").",
	},
	"G21": {
		Group: Groups[Units],
		Code:  "G21",
		Label: "Units will be in mm",
		Notes: "Best practice: do this at the start of a program and nowhere else. The usual minimum increment in G21 (one thousandth of a millimeter, .001 mm, that is, one micrometre).",
	},
	"M3": {
		Group: Groups[SpindleTurning],
		Code:  "M3",
		Label: "Spindle direction clockwise",
		Notes: "Starts or restarts the spindle spinning clockwise, if the system is wired up to start/stop the spindle.",
	},
	"M4": {
		Group: Groups[SpindleTurning],
		Code:  "M4",
		Label: "Spindle direction counter-clockwise",
		Notes: "Used to enable laser mode movement in Grbl 1.0 and later.",
	},
	"M5": {
		Group: Groups[SpindleTurning],
		Code:  "M5",
		Label: "Spindle direction clockwise",
		Notes: "Starts or restarts the spindle spinning clockwise, if the system is wired up to start/stop the spindle.",
	},
	"M7": {
		Group: Groups[Coolant],
		Code:  "M7",
		Label: "Mist",
		Notes: "Coolant control",
	},
	"M8": {
		Group: Groups[Coolant],
		Code:  "M8",
		Label: "Flood coolant on",
		Notes: "Coolant control",
	},
	"M9": {
		Group: Groups[Coolant],
		Code:  "M9",
		Label: "All coolant off.",
		Notes: "Coolant control",
	},
	"M6": {
		Group: Groups[ToolChange],
		Code:  "M6",
		Label: "Tool Change",
		Notes: "Coolant control",
	},
	"T?": {
		Group: Groups[Parameters],
		Code:  "T?",
		Label: "Tool Number",
	},
	"F?": {
		Group: Groups[Parameters],
		Code:  "F?",
		Label: "Feed Rate",
	},
	"S?": {
		Group: Groups[Parameters],
		Code:  "S?",
		Label: "Spindle Speed",
	},
	"X?": {
		Group: Groups[Parameters],
		Code:  "X?",
		Label: "X Axis Position",
		Notes: "",
	},
	"Y?": {
		Group: Groups[Parameters],
		Code:  "Y?",
		Label: "Y Axis Position",
		Notes: "",
	},
	"Z?": {
		Group: Groups[Parameters],
		Code:  "Z?",
		Label: "Z Axis Position",
		Notes: "",
	},
	"A?": {
		Group: Groups[Parameters],
		Code:  "A?",
		Label: "A Axis Position",
		Notes: "",
	},
	"B?": {
		Group: Groups[Parameters],
		Code:  "B?",
		Label: "B Axis Position",
		Notes: "",
	},
	"C?": {
		Group: Groups[Parameters],
		Code:  "C?",
		Label: "C Axis Position",
		Notes: "",
	},
	"I?": {
		Group: Groups[Parameters],
		Code:  "I?",
		Label: "Arc Centre in X Axis ",
		Notes: "",
	},
	"J?": {
		Group: Groups[Parameters],
		Code:  "J?",
		Label: "Arc Centre in Y Axis ",
		Notes: "",
	},
	"K?": {
		Group: Groups[Parameters],
		Code:  "K?",
		Label: "Arc Centre in Z Axis ",
		Notes: "",
	},
	"R?": {
		Group: Groups[Parameters],
		Code:  "R?",
		Label: "Arc Radius Size ",
		Notes: "",
	},
	"P?": {
		Group: Groups[Parameters],
		Code:  "P?",
		Label: "Parameter Address",
		Notes: "",
	},
}
