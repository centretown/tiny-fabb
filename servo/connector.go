// Copyright (c) 2021 Dave Marsh. See LICENSE.

package servo

import (
	"fmt"
	"html/template"
)

type Connector struct {
	Servos []*Servo
	layout *template.Template
}

func NewConnector(controlUrl string,
	layout *template.Template,
	count int) (conn *Connector) {

	conn = &Connector{
		Servos: make([]*Servo, count),
		layout: layout,
	}

	for i := range conn.Servos {
		conn.Servos[i] = &Servo{
			Title:      fmt.Sprintf("Servo-%02d", i),
			ControlUrl: controlUrl,
			Index:      uint(i),
			Command:    ServoEase,
			Angle:      90,
			Speed:      50,
			EaseType:   0,
		}
	}
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
