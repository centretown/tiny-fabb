package main

import (
	"fmt"
	"testing"

	"github.com/centretown/tiny-fabb/serialio"
)

func TestFindAvailableSerialPorts(t *testing.T) {
	s := serialio.ListSerial()
	for _, n := range s {
		fmt.Println(n)
	}
}
