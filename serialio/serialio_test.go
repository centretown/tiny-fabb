package serialio

import (
	"testing"
)

func TestSerialIo(t *testing.T) {
	// var id string
	for _, l := range ListSerial() {
		t.Log(l)
	}
}

// func testSerialIo(t *testing.T) {
// 	var id string
// 	for _, l := range ListSerial() {
// 		t.Log(l)
// 		if strings.Index(l, "ttyACM") != -1 {
// 			id = l
// 		}
// 	}

// 	sio, err := GetSerialIO(id)

// 	rdr, wrt, err := sio.Open()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	var (
// 		bufReader = bufio.NewReaderSize(rdr, 256)
// 		bufWriter = bufio.NewWriterSize(wrt, 64)
// 		line      []byte
// 	)

// 	_, err = bufWriter.WriteString("$$" + "\r")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	bufWriter.Flush()

// 	for {
// 		line, err = bufReader.ReadBytes('\n')
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			t.Fatal(err)
// 		}

// 		s := strings.TrimSpace(string(line))
// 		t.Log(s)
// 		if strings.HasPrefix(s, "ok") {
// 			break
// 		}
// 		if strings.HasPrefix(s, "error") {
// 			break
// 		}
// 	}

// 	err = sio.Close()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
