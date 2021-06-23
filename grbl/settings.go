package grbl

// GrblSettings as defined by grbl
type GrblSettings struct {
	StepPulse          uint    `json:"stepPulse"`
	StepIdleDelay      uint    `json:"stepIdleDelay"`
	StepPortInvertMask uint    `json:"stepPortInvertMask"`
	DirPortInvertMask  uint    `json:"dirPortInvertMask"`
	StepEnableInvert   bool    `json:"stepEnableInvert"`
	LimitPinsInvert    bool    `json:"limitPinsInvert"`
	ProbePinInvert     bool    `json:"probePinInvert"`
	StatusReportMask   uint    `json:"statusReportMask"`
	JunctionDeviation  float32 `json:"junctionDeviation"`
	ArcTolerance       float32 `json:"arcTolerance"`
	ReportInches       bool    `json:"reportInches"`
	SoftLimits         bool    `json:"softLimits"`
	HardLimits         bool    `json:"hardLimits"`
	HomingCycle        bool    `json:"homingCycle"`
	HomingDirInvert    uint    `json:"homingDirInvert"`
	HomingFeed         float32 `json:"homingFeed"`
	HomingSeek         float32 `json:"homingSeek"`
	HomingDebounce     uint    `json:"homingDebounce"`
	HomingPulloff      float32 `json:"homingPulloff"`
	MaxSpindleSpeed    float32 `json:"maxSpindleSpeed"`
	MinSpindleSpeed    float32 `json:"minSpindleSpeed"`
	LaserMode          bool    `json:"laserMode"`
	StepsX             float32 `json:"stepsX"`
	StepsY             float32 `json:"stepsY"`
	StepsZ             float32 `json:"stepsZ"`
	StepsA             float32 `json:"stepsA"`
	StepsB             float32 `json:"stepsB"`
	StepsC             float32 `json:"stepsC"`
	MaxRateX           float32 `json:"maxRateX"`
	MaxRateY           float32 `json:"maxRateY"`
	MaxRateZ           float32 `json:"maxRateZ"`
	MaxRateA           float32 `json:"maxRateA"`
	MaxRateB           float32 `json:"maxRateB"`
	MaxRateC           float32 `json:"maxRateC"`
	AccelX             float32 `json:"accelX"`
	AccelY             float32 `json:"accelY"`
	AccelZ             float32 `json:"accelZ"`
	AccelA             float32 `json:"accelA"`
	AccelB             float32 `json:"accelB"`
	AccelC             float32 `json:"accelC"`
	MaxTravelX         float32 `json:"maxTravelX"`
	MaxTravelY         float32 `json:"maxTravelY"`
	MaxTravelZ         float32 `json:"maxTravelZ"`
	MaxTravelA         float32 `json:"maxTravelA"`
	MaxTravelB         float32 `json:"maxTravelB"`
	MaxTravelC         float32 `json:"maxTravelC"`
}
