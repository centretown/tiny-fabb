// Copyright (c) 2021 Dave Marsh. See LICENSE.

package servo

import "github.com/centretown/tiny-fabb/forms"

var commandEntries = map[forms.WebId]forms.Entries{
	idIndex: {
		{
			ID:    idIndex.String(),
			Code:  "index",
			Label: "Servo",
			Type:  "readonly",
		},
	},
	idCommand: {
		{
			ID:    idCommand.String(),
			Code:  "command",
			Label: "Command",
			Type:  "number",
			Min:   1,
			Max:   5,
			Step:  1,
		},
	},
	idAngle: {
		{
			ID:    idAngle.String(),
			Code:  "angle",
			Label: "Angle",
			Type:  "number",
			Min:   0,
			Max:   180,
			Step:  1,
		},
	},
	idSpeed: {
		{
			ID:    idSpeed.String(),
			Code:  "speed",
			Label: "Speed",
			Type:  "number",
			Min:   1,
			Max:   255,
			Step:  1,
		},
	},
	idEaseType: {
		{
			ID:    idEaseType.String(),
			Code:  "type",
			Label: "Ease Type",
			Type:  "number",
			Min:   0,
			Max:   255,
			Step:  1,
		},
	},
}

var settingsEntries = map[forms.WebId]forms.Entries{
	idStep: {
		{
			ID:    idStep.String(),
			Label: "Step Size",
			Type:  "number",
			Min:   1,
			Max:   180,
			Step:  1,
		},
	},
	idMin: {
		{
			ID:    idMin.String(),
			Label: "Minimum",
			Type:  "number",
			Min:   0,
			Max:   179,
			Step:  1,
		},
	},
	idMax: {
		{
			ID:    idMax.String(),
			Label: "Maximum",
			Type:  "number",
			Min:   1,
			Max:   180,
			Step:  1,
		},
	},
	idInvert: {
		{
			ID:    idInvert.String(),
			Label: "Invert Direction",
			Type:  "checkbox",
		},
	},
}
