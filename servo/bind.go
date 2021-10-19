// Copyright (c) 2021 Dave Marsh. See LICENSE.

package servo

import "github.com/centretown/tiny-fabb/forms"

func (svo *Servo) bind() {
	svo.CommandForms = forms.Forms{
		idCommand: {
			ID:      idCommand,
			Value:   &svo.Command,
			Entries: commandEntries[idCommand],
		},
		idIndex: {
			ID:      idIndex,
			Value:   &svo.Index,
			Entries: commandEntries[idIndex],
		},
		idAngle: {
			ID:      idAngle,
			Value:   &svo.Angle,
			Entries: commandEntries[idAngle],
		},
		idSpeed: {
			ID:      idSpeed,
			Value:   &svo.Speed,
			Entries: commandEntries[idSpeed],
		},
		idEaseType: {
			ID:      idEaseType,
			Value:   &svo.EaseType,
			Entries: commandEntries[idEaseType],
		},
	}

	svo.SettingForms = forms.Forms{
		idStep: {
			ID:      idStep,
			Value:   &svo.Settings.Step,
			Entries: settingsEntries[idStep],
		},
		idMin: {
			ID:      idMin,
			Value:   &svo.Settings.Min,
			Entries: settingsEntries[idMin],
		},
		idMax: {
			ID:      idMax,
			Value:   &svo.Settings.Max,
			Entries: settingsEntries[idMax],
		},
		idInvert: {
			ID:      idInvert,
			Value:   &svo.Settings.Invert,
			Entries: settingsEntries[idInvert],
		},
	}
}
