// Copyright (c) 2021 Dave Marsh. See LICENSE.

package camera

type Servo struct {
	ControlUrl string `json:"controlUrl"`
	Index      uint   `json:"index"`

	Command  uint `json:"command"`
	Angle    uint `json:"angle"`
	Speed    uint `json:"speed"`
	EaseType uint `json:"easeType"`
}
