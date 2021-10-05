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

func ServoCommands() []ServoCommand {
	return []ServoCommand{ServoHome, ServoMove, ServoEase}
}

type Servo struct {
	Title      string       `json:"title"`
	ControlUrl string       `json:"controlUrl"`
	Index      uint         `json:"index"`
	Command    ServoCommand `json:"command"`
	Angle      uint         `json:"angle"`
	Speed      uint         `json:"speed"`
	EaseType   uint         `json:"easeType"`
	Forms      forms.Forms  `json:"-"`
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
	_, err = forms.Request(b.String())
	if err != nil {
		forms.WriteError(w, err)
	}
}

func (svo *Servo) Home(w http.ResponseWriter, r *http.Request) {
	svo.Command = ServoHome
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

func (svo *Servo) Connect(router *mux.Router, prefix string) {
	svo.bind()
	index := fmt.Sprintf("/servo%d", svo.Index)
	router.HandleFunc(index+"/home/{angle}/{speed}/{easeType}/", svo.Home)
	router.HandleFunc(index+"/home/", svo.Home)
	router.HandleFunc(index+"/move/{angle}/{speed}/{easeType}/", svo.Move)
	router.HandleFunc(index+"/move/{angle}/{speed}/", svo.Move)
	router.HandleFunc(index+"/ease/{angle}/{speed}/{easeType}/", svo.Ease)
}
