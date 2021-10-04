// Copyright (c) 2021 Dave Marsh. See LICENSE.

package servo

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/centretown/tiny-fabb/forms"
	"github.com/gorilla/mux"
)

type ServoCommand uint

const (
	ServoBase ServoCommand = iota
	ServoHome
	ServoMove
	ServoEase
	ServoMax
)

var svoCommandText = []string{
	"invalid",
	"home",
	"move",
	"ease",
}

func (svoc ServoCommand) String() string {
	i := svoc
	if i >= ServoMax {
		i = ServoBase
	}
	return svoCommandText[i]
}

type Servo struct {
	ControlUrl string       `json:"controlUrl"`
	Index      uint         `json:"index"`
	Command    ServoCommand `json:"command"`
	Angle      uint         `json:"angle"`
	Speed      uint         `json:"speed"`
	EaseType   uint         `json:"easeType"`
	Forms      forms.Forms  `json:"-"`
}

func (svo *Servo) apply(w http.ResponseWriter) {
	// _, err = request()
	// if err != nil {
	// 	return
	// }
	b := &strings.Builder{}
	_, err := fmt.Fprintf(b,
		"%s?var=%s&index=%d&angle=%d&speed=%d&type=%d",
		svo.ControlUrl, svo.Command.String(),
		svo.Index, svo.Angle, svo.Speed, svo.EaseType)

	if err != nil {
		forms.WriteError(w, err)
		return
	}
	fmt.Fprint(w, b.String())
}

func (svo *Servo) Connect(router *mux.Router, prefix string) {
	svo.bind()
	index := fmt.Sprintf("%sservo%d", prefix, svo.Index)
	router.HandleFunc(index+"/home/", func(w http.ResponseWriter, r *http.Request) {
		svo.Command = ServoHome
		svo.Angle = 0
		svo.Speed = 0
		svo.EaseType = 0
		svo.apply(w)
	})

	router.HandleFunc(index+"/move/{angle}/", func(w http.ResponseWriter, r *http.Request) {
		svo.Command = ServoMove
		svo.Angle = forms.GetRequestUint(r, "angle")
		svo.Speed = 50
		svo.EaseType = 0
		svo.apply(w)
	})
	router.HandleFunc(index+"/move/{angle}/{speed}/", func(w http.ResponseWriter, r *http.Request) {
		svo.Command = ServoMove
		svo.Angle = forms.GetRequestUint(r, "angle")
		svo.Speed = forms.GetRequestUint(r, "speed")
		svo.EaseType = 0
		svo.apply(w)
	})

	router.HandleFunc(index+"/ease/{angle}/", func(w http.ResponseWriter, r *http.Request) {
		svo.Command = ServoEase
		svo.Angle = forms.GetRequestUint(r, "angle")
		svo.Speed = 50
		svo.EaseType = 0
		svo.apply(w)
	})
	router.HandleFunc(index+"/ease/{angle}/{speed}/{easeType}/", func(w http.ResponseWriter, r *http.Request) {
		svo.Command = ServoEase
		svo.Angle = forms.GetRequestUint(r, "angle")
		svo.Speed = forms.GetRequestUint(r, "speed")
		svo.EaseType = forms.GetRequestUint(r, "easeType")
		svo.apply(w)
	})
}
