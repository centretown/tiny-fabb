package main

import (
	"bufio"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/tarm/serial"
)

// func TestSerial(t *testing.T) {
func estSerial(t *testing.T) {
	config := &serial.Config{
		Name:     "/dev/ttyACM0",
		Baud:     115200,
		Size:     8,
		Parity:   'N',
		StopBits: 1,
	}
	// /dev/ttyACM0  115200,8,N,1
	serialPort, err := serial.OpenPort(config)
	if err != nil {
		t.Logf("Open Serial port:%v error:%v\n", config.Name, err)
		return
	}

	defer serialPort.Close()

	var (
		in  = make(chan string, 10)
		out = make(chan string)
	)

	go testMonitor(in, out, serialPort)
	go capture(in)

	// out <- "$$"
	// out <- "$I"
	// // out <- "?"
	// // out <- "$G"
	// // out <- "$N"

	// out <- "$C"
	// // out <- "$X"
	// // out <- "$H"
	// out <- "$"
	out <- "G0 X20 Y30 Z10 F900"
	out <- "G0 X20 Y30 Z10 F900"
	time.Sleep(time.Second * 2)
}

func capture(in <-chan string) (s string) {

	for {
		select {
		case s = <-in:
			fmt.Print(s)
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

func testMonitor(in chan<- string, out <-chan string, port *serial.Port) {
	var (
		bufReader = bufio.NewReaderSize(port, 4096)
		bufWriter = bufio.NewWriterSize(port, 64)
		line      []byte
		err       error
	)

	for {
		select {
		case tx := <-out:
			_, err = bufWriter.WriteString(tx + "\r")
			if err != nil {
				fmt.Println(err)
			}
			bufWriter.Flush()
		default:
			line, err = bufReader.ReadSlice('\n')
			if err != nil && err != io.EOF {
				fmt.Println(err)
			} else {
				s := string(line)
				in <- s
			}
		}
		time.Sleep(time.Millisecond)
	}
}
