// Copyright (c) 2021 Dave Marsh. See LICENSE.

package servo

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
)

type Connector struct {
	Servos     []*Servo `json:"servos"`
	layout     *template.Template
	controlUrl string
}

func NewConnector(controlUrl string, dataSource string,
	layout *template.Template,
	count int) (conn *Connector) {

	conn = &Connector{
		Servos:     make([]*Servo, count),
		layout:     layout,
		controlUrl: controlUrl,
	}
	err := conn.Load(dataSource)
	if err != nil {
		conn.Initialize()
		conn.Save(dataSource)
	}
	return
}

func (conn *Connector) Initialize() {
	for i := range conn.Servos {
		conn.Servos[i] = &Servo{
			Title:      fmt.Sprintf("Servo-%02d", i),
			ControlUrl: conn.controlUrl,
			Index:      uint(i),
			Command:    ServoEase,
			Angle:      90,
			Speed:      50,
			EaseType:   0,
			Settings: Settings{
				Step:   18,
				Min:    0,
				Max:    180,
				Invert: false,
			},
		}
	}
}

func (conn *Connector) Save(dataSource string) (err error) {
	b, err := json.MarshalIndent(conn.Servos, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dataSource+"servos.json", b, 0640)
	return

}

func (conn *Connector) Load(dataSource string) (err error) {
	var b []byte
	b, err = ioutil.ReadFile(dataSource + "servos.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &conn.Servos)

	return
}

func (conn *Connector) Connect(indeces ...uint) (servos []*Servo, err error) {
	servos = make([]*Servo, len(indeces))
	l := uint(len(conn.Servos))
	for i, index := range indeces {
		if index >= l {
			err = fmt.Errorf("index (%d) out of range", index)
			return
		}
		servos[i] = conn.Servos[index]
	}
	return
}

func (conn *Connector) Count() int {
	return len(conn.Servos)
}
