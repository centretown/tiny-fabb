package grbl

import "github.com/centretown/tiny-fabb/forms"

const (
	settingsBaseURL = "https://github.com/gnea/grbl/wiki/Grbl-v1.1-Configuration"
	settingsURL     = "#grbl-settings"
	commandsBaseURL = "https://github.com/gnea/grbl/wiki/Grbl-v1.1-Commands"
	commandsURL     = "#grbl-commands"
)

func (g *Controller) bindSettings() (frms forms.Forms) {
	frms = forms.Forms{
		idStepPulse: {
			ID:      idStepPulse,
			Value:   &g.Settings.StepPulse,
			Entries: settingEntries[idStepPulse],
		},
		idStepIdleDelay: {
			ID:      idStepIdleDelay,
			Value:   &g.Settings.StepIdleDelay,
			Entries: settingEntries[idStepIdleDelay],
		},
		idStepPortInvertMask: {
			ID:      idStepPortInvertMask,
			Value:   &g.Settings.StepPortInvertMask,
			Entries: settingEntries[idStepPortInvertMask],
		},
		idDirPortInvertMask: {
			ID:      idDirPortInvertMask,
			Value:   &g.Settings.DirPortInvertMask,
			Entries: settingEntries[idDirPortInvertMask],
		},
		idStepEnableInvert: {
			ID:      idStepEnableInvert,
			Value:   &g.Settings.StepEnableInvert,
			Entries: settingEntries[idStepEnableInvert],
		},
		idLimitPinsInvert: {
			ID:      idLimitPinsInvert,
			Value:   &g.Settings.LimitPinsInvert,
			Entries: settingEntries[idLimitPinsInvert],
		},
		idProbePinInvert: {
			ID:      idProbePinInvert,
			Value:   &g.Settings.ProbePinInvert,
			Entries: settingEntries[idProbePinInvert],
		},
		idStatusReportMask: {
			ID:      idStatusReportMask,
			Value:   &g.Settings.StatusReportMask,
			Entries: settingEntries[idStatusReportMask],
		},
		idJunctionDeviation: {
			ID:      idJunctionDeviation,
			Value:   &g.Settings.JunctionDeviation,
			Entries: settingEntries[idJunctionDeviation],
		},
		idArcTolerance: {
			ID:      idArcTolerance,
			Value:   &g.Settings.ArcTolerance,
			Entries: settingEntries[idArcTolerance],
		},
		idReportInches: {
			ID:      idReportInches,
			Value:   &g.Settings.ReportInches,
			Entries: settingEntries[idReportInches],
		},
		idSoftLimits: {
			ID:      idSoftLimits,
			Value:   &g.Settings.SoftLimits,
			Entries: settingEntries[idSoftLimits],
		},
		idHardLimits: {
			ID:      idHardLimits,
			Value:   &g.Settings.HardLimits,
			Entries: settingEntries[idHardLimits],
		},
		idHomingCycle: {
			ID:      idHomingCycle,
			Value:   &g.Settings.HomingCycle,
			Entries: settingEntries[idHomingCycle],
		},
		idHomingDirInvert: {
			ID:      idHomingDirInvert,
			Value:   &g.Settings.HomingDirInvert,
			Entries: settingEntries[idHomingDirInvert],
		},
		idHomingFeed: {
			ID:      idHomingFeed,
			Value:   &g.Settings.HomingFeed,
			Entries: settingEntries[idHomingFeed],
		},
		idHomingSeek: {
			ID:      idHomingSeek,
			Value:   &g.Settings.HomingSeek,
			Entries: settingEntries[idHomingSeek],
		},
		idHomingDebounce: {
			ID:      idHomingDebounce,
			Value:   &g.Settings.HomingDebounce,
			Entries: settingEntries[idHomingDebounce],
		},
		idHomingPulloff: {
			ID:      idHomingPulloff,
			Value:   &g.Settings.HomingPulloff,
			Entries: settingEntries[idHomingPulloff],
		},
		idMaxSpindleSpeed: {
			ID:      idMaxSpindleSpeed,
			Value:   &g.Settings.MaxSpindleSpeed,
			Entries: settingEntries[idMaxSpindleSpeed],
		},
		idMinSpindleSpeed: {
			ID:      idMinSpindleSpeed,
			Value:   &g.Settings.MinSpindleSpeed,
			Entries: settingEntries[idMinSpindleSpeed],
		},
		idLaserMode: {
			ID:      idLaserMode,
			Value:   &g.Settings.LaserMode,
			Entries: settingEntries[idLaserMode],
		},
		idStepsX: {
			ID:      idStepsX,
			Value:   &g.Settings.StepsX,
			Entries: settingEntries[idStepsX],
		},
		idStepsY: {
			ID:      idStepsY,
			Value:   &g.Settings.StepsY,
			Entries: settingEntries[idStepsY],
		},
		idStepsZ: {
			ID:      idStepsZ,
			Value:   &g.Settings.StepsZ,
			Entries: settingEntries[idStepsZ],
		},
		idStepsA: {
			ID:      idStepsA,
			Value:   &g.Settings.StepsA,
			Entries: settingEntries[idStepsA],
		},
		idStepsB: {
			ID:      idStepsB,
			Value:   &g.Settings.StepsB,
			Entries: settingEntries[idStepsB],
		},
		idStepsC: {
			ID:      idStepsC,
			Value:   &g.Settings.StepsC,
			Entries: settingEntries[idStepsC],
		},
		idMaxRateX: {
			ID:      idMaxRateX,
			Value:   &g.Settings.MaxRateX,
			Entries: settingEntries[idMaxRateX],
		},
		idMaxRateY: {
			ID:      idMaxRateY,
			Value:   &g.Settings.MaxRateY,
			Entries: settingEntries[idMaxRateY],
		},
		idMaxRateZ: {
			ID:      idMaxRateZ,
			Value:   &g.Settings.MaxRateZ,
			Entries: settingEntries[idMaxRateZ],
		},
		idMaxRateA: {
			ID:      idMaxRateA,
			Value:   &g.Settings.MaxRateA,
			Entries: settingEntries[idMaxRateA],
		},
		idMaxRateB: {
			ID:      idMaxRateB,
			Value:   &g.Settings.MaxRateB,
			Entries: settingEntries[idMaxRateB],
		},
		idMaxRateC: {
			ID:      idMaxRateC,
			Value:   &g.Settings.MaxRateC,
			Entries: settingEntries[idMaxRateC],
		},
		idAccelX: {
			ID:      idAccelX,
			Value:   &g.Settings.AccelX,
			Entries: settingEntries[idAccelX],
		},
		idAccelY: {
			ID:      idAccelY,
			Value:   &g.Settings.AccelY,
			Entries: settingEntries[idAccelY],
		},
		idAccelZ: {
			ID:      idAccelZ,
			Value:   &g.Settings.AccelZ,
			Entries: settingEntries[idAccelZ],
		},
		idAccelA: {
			ID:      idAccelA,
			Value:   &g.Settings.AccelA,
			Entries: settingEntries[idAccelA],
		},
		idAccelB: {
			ID:      idAccelB,
			Value:   &g.Settings.AccelB,
			Entries: settingEntries[idAccelB],
		},
		idAccelC: {
			ID:      idAccelC,
			Value:   &g.Settings.AccelC,
			Entries: settingEntries[idAccelC],
		},
		idMaxTravelX: {
			ID:      idMaxTravelX,
			Value:   &g.Settings.MaxTravelX,
			Entries: settingEntries[idMaxTravelX],
		},
		idMaxTravelY: {
			ID:      idMaxTravelY,
			Value:   &g.Settings.MaxTravelY,
			Entries: settingEntries[idMaxTravelY],
		},
		idMaxTravelZ: {
			ID:      idMaxTravelZ,
			Value:   &g.Settings.MaxTravelZ,
			Entries: settingEntries[idMaxTravelZ],
		},
		idMaxTravelA: {
			ID:      idMaxTravelA,
			Value:   &g.Settings.MaxTravelA,
			Entries: settingEntries[idMaxTravelA],
		},
		idMaxTravelB: {
			ID:      idMaxTravelB,
			Value:   &g.Settings.MaxTravelB,
			Entries: settingEntries[idMaxTravelB],
		},
		idMaxTravelC: {
			ID:      idMaxTravelC,
			Value:   &g.Settings.MaxTravelC,
			Entries: settingEntries[idMaxTravelC],
		},
		idSettings: {
			ID:      idSettings,
			Entries: commandEntries[idSettings],
			Value:   g.Commands.Settings,
		},
	}
	for _, form := range frms {
		form.BaseUrl = settingsBaseURL
	}
	return
}

func (g *Controller) bindCommands() (frms forms.Forms) {
	frms = forms.Forms{
		idParameters: {
			ID:      idParameters,
			Entries: commandEntries[idParameters],
			Value:   g.Commands.Parameters,
		},
		idParserState: {
			ID:      idParserState,
			Entries: commandEntries[idParserState],
			Value:   g.Commands.ParserState,
		},
		idBuildInfo: {
			ID:      idBuildInfo,
			Entries: commandEntries[idBuildInfo],
			Value:   g.Commands.BuildInfo,
		},
		idStartupBlocks: {
			ID:      idStartupBlocks,
			Entries: commandEntries[idStartupBlocks],
			Value:   g.Commands.StartupBlocks,
		},
		idCodeMode: {
			ID:      idCodeMode,
			Entries: commandEntries[idCodeMode],
			Value:   g.Commands.CodeMode,
		},
		idKillAlarm: {
			ID:      idKillAlarm,
			Entries: commandEntries[idKillAlarm],
			Value:   g.Commands.KillAlarm,
		},
		idRunHomingCycle: {
			ID:      idRunHomingCycle,
			Entries: commandEntries[idRunHomingCycle],
			Value:   g.Commands.RunHomingCycle,
		},
		idRunJoggingCycle: {
			ID:      idRunJoggingCycle,
			Entries: commandEntries[idRunJoggingCycle],
			Value:   g.Commands.RunJoggingCycle,
		},
		idEraseRestore: {
			ID:      idEraseRestore,
			Entries: commandEntries[idEraseRestore],
			Value:   g.Commands.EraseRestore,
		},
		idEraseZero: {
			ID:      idEraseZero,
			Entries: commandEntries[idEraseZero],
			Value:   g.Commands.EraseZero,
		},
		idClearRestore: {
			ID:      idClearRestore,
			Entries: commandEntries[idClearRestore],
			Value:   g.Commands.ClearRestore,
		},
	}
	for _, form := range frms {
		form.BaseUrl = commandsBaseURL
	}
	return
}
