// Copyright (c) 2021 Dave Marsh. See LICENSE.

package servo

// {{$dY:=15}}{{$dX:=15}}{{$maxX:=180}}{{$maxY:=180}}{{$orgX:=0}}{{$orgY:=0}}{{$speed:=60}}
type Settings struct {
	Step   int  `json:"step"`
	Min    int  `json:"min"`
	Max    int  `json:"max"`
	Invert bool `json:"invert"`
}

type Position struct {
	Max  int
	Step int
}

func (set *Settings) Calc(stepCount int) (pos *Position) {
	pos = &Position{}
	pos.Step = (set.Max - set.Min) / stepCount
	if set.Invert {
		pos.Max = set.Min
		pos.Step *= -1
	} else {
		pos.Max = set.Max
	}
	return
}
