package grbl

import "github.com/centretown/tiny-fabb/web"

const (
	idBase               web.WebId = 0
	idBaseSettings       web.WebId = 0
	idStepPulse          web.WebId = 0
	idStepIdleDelay      web.WebId = 1
	idStepPortInvertMask web.WebId = 2
	idDirPortInvertMask  web.WebId = 3
	idStepEnableInvert   web.WebId = 4
	idLimitPinsInvert    web.WebId = 5
	idProbePinInvert     web.WebId = 6
	idStatusReportMask   web.WebId = 10
	idJunctionDeviation  web.WebId = 11
	idArcTolerance       web.WebId = 12
	idReportInches       web.WebId = 13
	idSoftLimits         web.WebId = 20
	idHardLimits         web.WebId = 21
	idHomingCycle        web.WebId = 22
	idHomingDirInvert    web.WebId = 23
	idHomingFeed         web.WebId = 24
	idHomingSeek         web.WebId = 25
	idHomingDebounce     web.WebId = 26
	idHomingPulloff      web.WebId = 27
	idMaxSpindleSpeed    web.WebId = 30
	idMinSpindleSpeed    web.WebId = 31
	idLaserMode          web.WebId = 32
	idStepsX             web.WebId = 100
	idStepsY             web.WebId = 101
	idStepsZ             web.WebId = 102
	idStepsA             web.WebId = 103
	idStepsB             web.WebId = 104
	idStepsC             web.WebId = 105
	idMaxRateX           web.WebId = 110
	idMaxRateY           web.WebId = 111
	idMaxRateZ           web.WebId = 112
	idMaxRateA           web.WebId = 113
	idMaxRateB           web.WebId = 114
	idMaxRateC           web.WebId = 115
	idAccelX             web.WebId = 120
	idAccelY             web.WebId = 121
	idAccelZ             web.WebId = 122
	idAccelA             web.WebId = 123
	idAccelB             web.WebId = 124
	idAccelC             web.WebId = 125
	idMaxTravelX         web.WebId = 130
	idMaxTravelY         web.WebId = 131
	idMaxTravelZ         web.WebId = 132
	idMaxTravelA         web.WebId = 133
	idMaxTravelB         web.WebId = 134
	idMaxTravelC         web.WebId = 135
	idSettings           web.WebId = 136
	idEndSettings        web.WebId = 137
)
const (
	idBaseCommands web.WebId = 2000 + iota
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
