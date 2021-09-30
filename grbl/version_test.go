// Copyright (c) 2021 Dave Marsh. See LICENSE.

package grbl

import (
	"testing"
)

func TestVersion(t *testing.T) {
	var testA = []string{
		"[VER:1.3a.20210424:]",
		"[OPT:PHBSW]",
		"[MSG:Using machine:Test Drive - Demo Only No I/O!]",
		"[MSG:Mode=STA:SSID=xlab:Status=Connected:IP=192.168.0.31:MAC=10-52-1C-67-83-44]",
		"[MSG:No BT]",
		"ok",
	}
	testResults(t, testA)

	var testB = []string{
		"[VER:1.1h.20190825:ID=1]",
		"[OPT:V,15,128]",
		"ok",
	}
	testResults(t, testB)

	var testC = []string{
		"[VER:1.3a.20210424:]",
		"[OPT:PHBSW]",
		"[MSG:Using machine:Test Drive - Demo Only No I/O!]",
		"[MSG:Mode=STA:SSID=xlab:Status=Connected:IP=192.168.0.72:MAC=9C-9C-1F-C4-B7-88]",
		"[MSG:No BT]",
		"ok",
	}
	testResults(t, testC)
}

func testResults(t *testing.T, results []string) {
	pr, err := CheckProfile(results)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pr.Print())
}
