// Copyright (c) 2021 Dave Marsh. See LICENSE.

package servo

import "github.com/centretown/tiny-fabb/forms"

const (
	idBase forms.WebId = iota + 6000
	idIndex
	idCommand
	idAngle
	idSpeed
	idEaseType
)
