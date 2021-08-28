package serialio

import (
	"fmt"
	"io"
	"time"

	"github.com/tarm/serial"
)

// SerialIO supports the MonitorIO interface for serial ports
type SerialIO struct {
	Name      string
	opened    bool
	available bool
	port      *serial.Port
}

// Open -
func (sio *SerialIO) Open(opts ...interface{}) (rdr io.Reader, wrt io.Writer, err error) {
	var (
		config *serial.Config
		ok     bool
	)

	if len(opts) > 0 {
		config, ok = opts[0].(*serial.Config)
	}

	if !ok {
		config = &serial.Config{
			Name:        sio.Name,
			Baud:        115200,
			Size:        8,
			Parity:      'N',
			StopBits:    1,
			ReadTimeout: time.Millisecond,
		}

	}

	sio.port, err = serial.OpenPort(config)
	if err != nil {
		err = fmt.Errorf("SerialIO: failed to open %v '%v'", sio.Name, err)
		return
	}
	rdr, wrt = sio.port, sio.port
	sio.opened = true
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
