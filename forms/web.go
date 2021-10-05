// Copyright (c) 2021 Dave Marsh. See LICENSE.

package forms

import (
	"fmt"
	"io"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

func GetRequestString(r *http.Request, key string) (sel string) {
	vars := mux.Vars(r)
	fmt.Sscan(vars[key], &sel)
	return
}

func GetRequestUint(r *http.Request, key string) (sel uint) {
	vars := mux.Vars(r)
	fmt.Sscan(vars[key], &sel)
	return
}

func GetRequestInt(r *http.Request, key string) (sel int) {
	vars := mux.Vars(r)
	fmt.Sscan(vars[key], &sel)
	return
}

func WriteError(w http.ResponseWriter, err error) {
	glog.Infoln(err)
	http.Error(w, err.Error(), 400)
}

func Request(u string) (data []byte, err error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	return
}
