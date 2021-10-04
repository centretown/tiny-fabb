// Copyright (c) 2021 Dave Marsh. See LICENSE.

package servo

import "github.com/centretown/tiny-fabb/forms"

var servoEntries = map[forms.WebId]forms.Entries{
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
			Max:   3,
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
			Max:   10,
			Step:  1,
		},
	},
}
