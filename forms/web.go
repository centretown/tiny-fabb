// Copyright (c) 2021 Dave Marsh. See LICENSE.

package forms

import (
	"fmt"
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

// type RequestVar struct {
// 	Code string
// 	Val  interface{}
// }

// func GetRequestVars(r *http.Request, reqVars []*RequestVar) (err error) {
// 	muxVars := mux.Vars(r)
// 	for _, v := range reqVars {
// 		val, ok := muxVars[v.Code]
// 		if !ok {
// 			err = fmt.Errorf("code '%s' not found", v.Code)
// 			return
// 		}
// 		_, err = fmt.Scan(val, v.Val)
// 		if err != nil {
// 			return
// 		}
// 	}
// 	return
// }
