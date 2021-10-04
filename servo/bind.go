// Copyright (c) 2021 Dave Marsh. See LICENSE.

package servo

import "github.com/centretown/tiny-fabb/forms"

func (svo *Servo) bind() {
	svo.Forms = forms.Forms{
		idCommand: {
			ID:      idCommand,
			Value:   &svo.Command,
			Entries: servoEntries[idCommand],
		},
		idIndex: {
			ID:      idIndex,
			Value:   &svo.Index,
			Entries: servoEntries[idIndex],
		},
		idAngle: {
			ID:      idAngle,
			Value:   &svo.Angle,
			Entries: servoEntries[idAngle],
		},
		idSpeed: {
			ID:      idSpeed,
			Value:   &svo.Speed,
			Entries: servoEntries[idSpeed],
		},
		idEaseType: {
			ID:      idEaseType,
			Value:   &svo.EaseType,
			Entries: servoEntries[idEaseType],
		},
	}
}
