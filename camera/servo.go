// Copyright (c) 2021 Dave Marsh. See LICENSE.

package camera

type ServoAction struct {
	Command  uint `json:"command"`
	Angle    uint `json:"angle"`
	Speed    uint `json:"speed"`
	EaseType uint `json:"easeType"`
}

type Servo struct {
	ControlUrl string         `json:"controlUrl"`
	Index      uint           `json:"index"`
	Queue      []*ServoAction `json:"-"`
}

func (svo *Servo) Add(act *ServoAction) {
	svo.Queue = append(svo.Queue, act)
}
