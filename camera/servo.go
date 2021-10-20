// Copyright (c) 2021 Dave Marsh. See LICENSE.

package camera

import (
	"net/http"

	"github.com/centretown/tiny-fabb/forms"
	"github.com/centretown/tiny-fabb/servo"
)

type ServoControl struct {
	ID       string
	Icon     string
	Speed    uint
	PanMax   int
	PanStep  int
	TiltMax  int
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
	"bi-arrow-repeat",
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
	)

	pan = cam.Servos[0]
	panPos = pan.Settings.Calc(nSteps)

	if count < 2 {
		start = servoLeft
		end = servoRight
	} else {
		tilt = cam.Servos[1]
		tiltPos = tilt.Settings.Calc(nSteps)
		end = servoDownRight
	}

	for i := start; i <= end; i++ {

		ctl := &ServoControl{
			ID:   cam.ID,
			Icon: servoIcons[i],
		}

		switch i {
		case servoUpLeft:
			ctl.Speed = (tilt.Speed + pan.Speed) / 2
			ctl.PanMax = -panPos.Max
			ctl.PanStep = -panPos.Step
			ctl.TiltMax = tiltPos.Max
			ctl.TiltStep = tiltPos.Step
		case servoUp:
			ctl.Speed = tilt.Speed
			ctl.PanMax = 0
			ctl.PanStep = 0
			ctl.TiltMax = tiltPos.Max
			ctl.TiltStep = tiltPos.Step
		case servoUpRight:
			ctl.Speed = (tilt.Speed + pan.Speed) / 2
			ctl.PanMax = panPos.Max
			ctl.PanStep = panPos.Step
			ctl.TiltMax = tiltPos.Max
			ctl.TiltStep = tiltPos.Step
		case servoLeft:
			ctl.Speed = pan.Speed
			ctl.PanMax = -panPos.Max
			ctl.PanStep = -panPos.Step
			ctl.TiltMax = 0
			ctl.TiltStep = 0
		case servoMiddle:
			ctl.Speed = 0
			ctl.PanMax = 0
			ctl.PanStep = 0
			ctl.TiltMax = 0
			ctl.TiltStep = 0
		case servoRight:
			ctl.Speed = pan.Speed
			ctl.PanMax = panPos.Max
			ctl.PanStep = panPos.Step
			ctl.TiltMax = 0
			ctl.TiltStep = 0
		case servoDownLeft:
			ctl.Speed = (tilt.Speed + pan.Speed) / 2
			ctl.PanMax = -panPos.Max
			ctl.PanStep = -panPos.Step
			ctl.TiltMax = -tiltPos.Max
			ctl.TiltStep = -tiltPos.Step
		case servoDown:
			ctl.Speed = tilt.Speed
			ctl.PanMax = 0
			ctl.PanStep = 0
			ctl.TiltMax = -tiltPos.Max
			ctl.TiltStep = -tiltPos.Step
		case servoDownRight:
			ctl.Speed = (tilt.Speed + pan.Speed) / 2
			ctl.PanMax = panPos.Max
			ctl.PanStep = panPos.Step
			ctl.TiltMax = -tiltPos.Max
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

	pan := forms.GetRequestInt(r, "pan")
	tilt := forms.GetRequestInt(r, "tilt")

	if pan != 0 {
		svo := cam.Servos[0]
		cam.move(w, svo, int(svo.Angle)+pan, speed)
	}

	if tilt == 0 || servoCount < 2 {
		w.Write([]byte("]"))
		return
	}

	if pan != 0 {
		w.Write([]byte(","))
	}

	svo := cam.Servos[1]
	cam.move(w, svo, int(svo.Angle)+tilt, speed)
	w.Write([]byte("]"))
}
