// Copyright (c) 2021 Dave Marsh. See LICENSE.

package camera

import (
	"fmt"
	"net/http"

	"github.com/centretown/tiny-fabb/forms"
	"github.com/centretown/tiny-fabb/servo"
)

type ServoControl struct {
	Icon     string
	Speed    uint
	PanID    string
	PanStep  int
	TiltID   string
	TiltStep int
}

const (
	servoUpLeft uint = iota
	servoUp
	servoUpRight
	servoLeft
	servoMiddle
	servoRight
	servoDownLeft
	servoDown
	servoDownRight
)

var servoIcons = []string{
	"bi-arrow-up-left",
	"bi-arrow-up",
	"bi-arrow-up-right",
	"bi-arrow-left",
	"bi-arrows-move",
	"bi-arrow-right",
	"bi-arrow-down-left",
	"bi-arrow-down",
	"bi-arrow-down-right",
}

func (cam *Camera) ServoControls(nSteps int) (ctls []*ServoControl) {
	ctls = make([]*ServoControl, 0)
	count := len(cam.Servos)
	if count < 1 {
		return
	}

	var (
		start, end      uint = 0, 0
		pan, tilt       *servo.Servo
		panPos, tiltPos *servo.Position
		panID, tiltID   string
	)

	pan = cam.Servos[0]
	panPos = pan.Settings.Calc(nSteps)
	panID = fmt.Sprintf("%s-%d", cam.ID, pan.Index)

	if count < 2 {
		start = servoLeft
		end = servoRight
	} else {
		tilt = cam.Servos[1]
		tiltPos = tilt.Settings.Calc(nSteps)
		tiltID = fmt.Sprintf("%s-%d", cam.ID, tilt.Index)
		end = servoDownRight
	}

	for i := start; i <= end; i++ {

		ctl := &ServoControl{
			Icon:   servoIcons[i],
			PanID:  panID,
			TiltID: tiltID,
		}

		switch i {
		case servoUpLeft:
			ctl.Speed = (tilt.Speed + pan.Speed) / 2
			ctl.PanStep = -panPos.Step
			ctl.TiltStep = tiltPos.Step
		case servoUp:
			ctl.Speed = tilt.Speed
			ctl.PanStep = 0
			ctl.TiltStep = tiltPos.Step
		case servoUpRight:
			ctl.Speed = (tilt.Speed + pan.Speed) / 2
			ctl.PanStep = panPos.Step
			ctl.TiltStep = tiltPos.Step
		case servoLeft:
			ctl.Speed = pan.Speed
			ctl.PanStep = -panPos.Step
			ctl.TiltStep = 0
		case servoMiddle:
			ctl.Speed = 0
			ctl.PanStep = 0
			ctl.TiltStep = 0
		case servoRight:
			ctl.Speed = pan.Speed
			ctl.PanStep = panPos.Step
			ctl.TiltStep = 0
		case servoDownLeft:
			ctl.Speed = (tilt.Speed + pan.Speed) / 2
			ctl.PanStep = -panPos.Step
			ctl.TiltStep = -tiltPos.Step
		case servoDown:
			ctl.Speed = tilt.Speed
			ctl.PanStep = 0
			ctl.TiltStep = -tiltPos.Step
		case servoDownRight:
			ctl.Speed = (tilt.Speed + pan.Speed) / 2
			ctl.PanStep = panPos.Step
			ctl.TiltStep = -tiltPos.Step
		}
		ctls = append(ctls, ctl)
	}

	return
}

func (cam *Camera) move(w http.ResponseWriter,
	svo *servo.Servo, angle int, speed uint) {

	if angle < svo.Settings.Min {
		angle = svo.Settings.Min
	} else if angle > svo.Settings.Max {
		angle = svo.Settings.Max
	}
	svo.Command = servo.ServoMove
	svo.Angle = uint(angle)
	svo.Speed = speed
	svo.Apply(w)
}

func (cam *Camera) PanTilt(w http.ResponseWriter, r *http.Request) {
	servoCount := len(cam.Servos)
	if servoCount < 1 {
		return
	}

	speed := forms.GetRequestUint(r, "speed")
	if speed < 1 || speed > 255 {
		speed = 50
	}

	w.Write([]byte("["))

	pan := forms.GetRequestUint(r, "pan")
	tilt := forms.GetRequestUint(r, "tilt")

	svo := cam.Servos[0]
	movePan := svo.Angle != pan
	if movePan {
		cam.move(w, svo, int(pan), speed)
	}

	if servoCount > 1 {
		svo = cam.Servos[1]
	}
	moveTilt := servoCount > 1 && svo.Angle != tilt
	if !moveTilt {
		w.Write([]byte("]"))
		return
	}

	if movePan {
		w.Write([]byte(","))
	}

	cam.move(w, svo, int(tilt), speed)
	w.Write([]byte("]"))
}
