// Copyright (c) 2021 Dave Marsh. See LICENSE.

package servo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestServo(t *testing.T) {
	conn := NewConnector("192.168.0.44", nil, 9)
	showServos(t, conn.Servos)

	svos, err := conn.Connect(0, 1, 4)
	if err != nil {
		t.Fatal(err)
	}
	showServos(t, svos)
	_, err = conn.Connect(0, 1, 9)
	if err == nil {
		t.Fatal()
	}
	t.Log(err)
	svos, err = conn.Connect(2, 3, 5, 6, 7, 8)
	if err != nil {
		t.Fatal(err)
	}
	showServos(t, svos)
}

func TestServoConnect(t *testing.T) {
	conn := NewConnector("http://192.168.0.44", nil, 9)
	svos, err := conn.Connect(0, 1)
	if err != nil {
		t.Fatal(err)
	}
	showServos(t, svos)

	router := mux.NewRouter()
	for _, svo := range svos {
		svo.Connect(router, "/camera0/")
	}

	router.Walk(func(route *mux.Route,
		router *mux.Router,
		ancestors []*mux.Route) (err error) {
		tmpl, err := route.GetPathTemplate()
		if err != nil {
			t.Log(err)
		}
		t.Log(tmpl)
		return
	})

	srvr := httptest.NewServer(router)
	defer srvr.Close()

	t.Log(srvr.URL)

	testRequest(t, srvr.URL+"/camera0/servo0/move/90/100/")
	testRequest(t, srvr.URL+"/camera0/servo1/ease/90/100/3/")
	testRequest(t, srvr.URL+"/camera0/servo1/ease/90/")
}

func showServos(t *testing.T, servos []*Servo) {
	b, err := json.MarshalIndent(servos, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}

func testRequest(t *testing.T, u string) {
	res, err := http.Get(u)
	if err != nil {
		t.Fatal(err)
		return
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(out))
}
