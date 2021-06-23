package grbl

import (
	"fmt"
	"testing"
)

func testConstants(t *testing.T) {
	fmt.Println("idStepPulse=", idStepPulse)
	fmt.Println("idStepIdleDelay=", idStepIdleDelay)
	fmt.Println("idStepPortInvertMask=", idStepPortInvertMask)
	fmt.Println("idDirPortInvertMask=", idDirPortInvertMask)
	fmt.Println("idStepEnableInvert=", idStepEnableInvert)
	fmt.Println("idLimitPinsInvert=", idLimitPinsInvert)
	fmt.Println("idProbePinInvert=", idProbePinInvert)

	fmt.Println("idStatusReportMask=", idStatusReportMask)
	fmt.Println("idJunctionDeviation=", idJunctionDeviation)
	fmt.Println("idArcTolerance=", idArcTolerance)
	fmt.Println("idReportInches=", idReportInches)

	fmt.Println("idSoftLimits=", idSoftLimits)
	fmt.Println("idHardLimits=", idHardLimits)
	fmt.Println("idHomingCycle=", idHomingCycle)
	fmt.Println("idHomingDirInvert=", idHomingDirInvert)
	fmt.Println("idHomingFeed=", idHomingFeed)
	fmt.Println("idHomingSeek=", idHomingSeek)
	fmt.Println("idHomingDebounce=", idHomingDebounce)
	fmt.Println("idHomingPulloff=", idHomingPulloff)

	fmt.Println("idMaxSpindleSpeed=", idMaxSpindleSpeed)
	fmt.Println("idMinSpindleSpeed=", idMinSpindleSpeed)
	fmt.Println("idLaserMode=", idLaserMode)

	fmt.Println("idStepsX=", idStepsX)
	fmt.Println("idStepsY=", idStepsY)
	fmt.Println("idStepsZ=", idStepsZ)

	fmt.Println("idMaxRateX=", idMaxRateX)
	fmt.Println("idMaxRateY=", idMaxRateY)
	fmt.Println("idMaxRateZ=", idMaxRateZ)

	fmt.Println("idAccelX=", idAccelX)
	fmt.Println("idAccelY=", idAccelY)
	fmt.Println("idAccelZ=", idAccelZ)

	fmt.Println("idMaxTravelX=", idMaxTravelX)
	fmt.Println("idMaxTravelY=", idMaxTravelY)
	fmt.Println("idMaxTravelZ=", idMaxTravelZ)
}
