package serialio

import (
	"fmt"
	"io"
	"time"

	"github.com/tarm/serial"
)

// SerialIO supports the MonitorIO interface for serial ports
type SerialIO struct {
	Name   string
	port   *serial.Port
	opened bool
}

var (
	availableSerialIO = make(map[string]*SerialIO)
)

// GetSerialIO -
func GetSerialIO(name string) (sio *SerialIO, err error) {
	sio, ok := availableSerialIO[name]
	if !ok {
		err = fmt.Errorf("%v not found", name)
	}
	return
}

// Open -
func (sio *SerialIO) Open() (rdr io.Reader, wrt io.Writer, err error) {
	config := &serial.Config{
		Name:        sio.Name,
		Baud:        115200,
		Size:        8,
		Parity:      'N',
		StopBits:    1,
		ReadTimeout: time.Millisecond,
	}

	sio.port, err = serial.OpenPort(config)
	if err != nil {
		err = fmt.Errorf("SerialIO: failed to open %v '%v'", sio.Name, err)
		return
	}
	rdr, wrt = sio.port, sio.port
	sio.opened = true
	availableSerialIO[sio.Name] = sio
	return
}

// Close -
func (sio *SerialIO) Close() (err error) {
	if sio.port == nil {
		err = fmt.Errorf("SerialIO: attempting to close nil port")
	} else {
		err = sio.port.Close()
		sio.port = nil
	}
	return
}

// ListSerialIO returns a the list enumerated ports
func ListSerialIO() (list []*SerialIO) {
	EnumerateSerialPorts()
	list = make([]*SerialIO, 0, len(availableSerialIO))
	for _, sio := range availableSerialIO {
		list = append(list, sio)
	}
	return
}

// ListSerial returns a the list enumerated ports
func ListSerial() (list []string) {
	EnumerateSerialPorts()
	list = make([]string, 0, len(availableSerialIO))
	for _, sio := range availableSerialIO {
		list = append(list, sio.Name)
	}
	return
}

const (
	sys = "/sys/class/tty/"
	dev = "/dev/"
	acm = "ttyACM"
	usb = "ttyUSB"
)

var filter = []string{acm, usb}

func includedPort(fname string) bool {
	for _, f := range filter {
		if strings.Contains(fname, f) {
			return true
		}
	}
	return false
}

func EnumerateSerialPorts() {
	var (
		files []os.FileInfo
		err   error
		name  string
	)

	files, err = ioutil.ReadDir(sys)
	if err != nil {
		return
	}

	for _, f := range files {
		name = f.Name()
		if includedPort(name) {
			name = dev + name
			availableSerialIO[name] = &SerialIO{Name: name}
		}
	}
}

func MonitorPorts() {
	for {
		time.Sleep(time.Minute)
		EnumerateSerialPorts()
	}
}
