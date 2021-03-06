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
	ServoTest
	ServoStop
	ServoMax
)

var svoCommandText = []string{
	"invalid",
	"home",
	"move",
	"ease",
	"test",
	"stop",
}

func (svoc ServoCommand) String() string {
	i := svoc
	if i >= ServoMax {
		i = ServoBase
	}
	return svoCommandText[i]
}

func ServoCommands() []ServoCommand {
	return []ServoCommand{ServoHome, ServoMove, ServoEase}
}

type Servo struct {
	Index        uint         `json:"index"`
	Title        string       `json:"title"`
	ControlUrl   string       `json:"controlUrl"`
	Settings     Settings     `json:"settings"`
	Command      ServoCommand `json:"command"`
	Angle        uint         `json:"angle"`
	Speed        uint         `json:"speed"`
	EaseType     uint         `json:"easeType"`
	SettingForms forms.Forms  `json:"-"`
	CommandForms forms.Forms  `json:"-"`
}

func (svo *Servo) Apply(w http.ResponseWriter) {
	b := &strings.Builder{}
	_, err := fmt.Fprintf(b,
		"%s?var=%s&index=%d&angle=%d&speed=%d&type=%d",
		svo.ControlUrl, svo.Command.String(),
		svo.Index, svo.Angle, svo.Speed, svo.EaseType)

	if err != nil {
		forms.WriteError(w, err)
		return
	}
	var data []byte
	data, err = forms.Request(b.String())
	if err != nil {
		forms.WriteError(w, err)
	}
	forms.WriteResponse(w, data)
}

func (svo *Servo) Home(w http.ResponseWriter, r *http.Request) {
	svo.Command = ServoHome
	svo.Angle = 0
	svo.Speed = 0
	svo.EaseType = 0
	svo.Apply(w)
}
func (svo *Servo) Stop(w http.ResponseWriter, r *http.Request) {
	svo.Command = ServoStop
	svo.Angle = 0
	svo.Speed = 0
	svo.EaseType = 0
	svo.Apply(w)
}
func (svo *Servo) Move(w http.ResponseWriter, r *http.Request) {
	svo.Command = ServoMove
	svo.Angle = forms.GetRequestUint(r, "angle")
	svo.Speed = forms.GetRequestUint(r, "speed")
	svo.EaseType = 0
	svo.Apply(w)
}

func (svo *Servo) Ease(w http.ResponseWriter, r *http.Request) {
	svo.Command = ServoEase
	svo.Angle = forms.GetRequestUint(r, "angle")
	svo.Speed = forms.GetRequestUint(r, "speed")
	svo.EaseType = forms.GetRequestUint(r, "easeType")
	svo.Apply(w)
}

func (svo *Servo) Test(w http.ResponseWriter, r *http.Request) {
	svo.Command = ServoTest
	svo.Angle = forms.GetRequestUint(r, "angle")
	svo.Speed = forms.GetRequestUint(r, "speed")
	svo.EaseType = forms.GetRequestUint(r, "easeType")
	svo.Apply(w)
}

func (svo *Servo) Connect(router *mux.Router, prefix string) {
	svo.bind()
	index := fmt.Sprintf("/servo%d", svo.Index)
	router.HandleFunc(index+"/home/{angle}/{speed}/{easeType}/", svo.Home)
	router.HandleFunc(index+"/home/", svo.Home)
	router.HandleFunc(index+"/stop/{angle}/{speed}/{easeType}/", svo.Stop)
	router.HandleFunc(index+"/stop/", svo.Stop)
	router.HandleFunc(index+"/move/{angle}/{speed}/{easeType}/", svo.Move)
	router.HandleFunc(index+"/move/{angle}/{speed}/", svo.Move)
	router.HandleFunc(index+"/ease/{angle}/{speed}/{easeType}/", svo.Ease)
	router.HandleFunc(index+"/test/{angle}/{speed}/{easeType}/", svo.Test)
}
