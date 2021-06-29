package grbl

import "github.com/centretown/tiny-fabb/forms"

const (
	idBase               forms.WebId = 0
	idBaseSettings       forms.WebId = 0
	idStepPulse          forms.WebId = 0
	idStepIdleDelay      forms.WebId = 1
	idStepPortInvertMask forms.WebId = 2
	idDirPortInvertMask  forms.WebId = 3
	idStepEnableInvert   forms.WebId = 4
	idLimitPinsInvert    forms.WebId = 5
	idProbePinInvert     forms.WebId = 6
	idStatusReportMask   forms.WebId = 10
	idJunctionDeviation  forms.WebId = 11
	idArcTolerance       forms.WebId = 12
	idReportInches       forms.WebId = 13
	idSoftLimits         forms.WebId = 20
	idHardLimits         forms.WebId = 21
	idHomingCycle        forms.WebId = 22
	idHomingDirInvert    forms.WebId = 23
	idHomingFeed         forms.WebId = 24
	idHomingSeek         forms.WebId = 25
	idHomingDebounce     forms.WebId = 26
	idHomingPulloff      forms.WebId = 27
	idMaxSpindleSpeed    forms.WebId = 30
	idMinSpindleSpeed    forms.WebId = 31
	idLaserMode          forms.WebId = 32
	idStepsX             forms.WebId = 100
	idStepsY             forms.WebId = 101
	idStepsZ             forms.WebId = 102
	idStepsA             forms.WebId = 103
	idStepsB             forms.WebId = 104
	idStepsC             forms.WebId = 105
	idMaxRateX           forms.WebId = 110
	idMaxRateY           forms.WebId = 111
	idMaxRateZ           forms.WebId = 112
	idMaxRateA           forms.WebId = 113
	idMaxRateB           forms.WebId = 114
	idMaxRateC           forms.WebId = 115
	idAccelX             forms.WebId = 120
	idAccelY             forms.WebId = 121
	idAccelZ             forms.WebId = 122
	idAccelA             forms.WebId = 123
	idAccelB             forms.WebId = 124
	idAccelC             forms.WebId = 125
	idMaxTravelX         forms.WebId = 130
	idMaxTravelY         forms.WebId = 131
	idMaxTravelZ         forms.WebId = 132
	idMaxTravelA         forms.WebId = 133
	idMaxTravelB         forms.WebId = 134
	idMaxTravelC         forms.WebId = 135
	idSettings           forms.WebId = 136
	idEndSettings        forms.WebId = 137
)
const (
	idBaseCommands forms.WebId = 2000 + iota
	idParameters
	idParserState
	idBuildInfo
	idStartupBlocks
	idCodeMode
	idKillAlarm
	idRunHomingCycle
	idRunJoggingCycle
	idEraseRestore
	idEraseZero
	idClearRestore
	idEndCommands
)
