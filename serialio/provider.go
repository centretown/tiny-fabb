package serialio

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/centretown/tiny-fabb/monitor"
)

const (
	sys = "/sys/class/tty/"
	dev = "/dev/"
	acm = "ttyACM"
	usb = "ttyUSB"
)

var defaultFilters = []string{acm, usb}

type Provider struct {
	Ports  map[string]*SerialIO
	Filter []string
}

func (prv *Provider) List() (list []string) {
	list = make([]string, 0)
	for _, p := range prv.Ports {
		if p.available {
			list = append(list, p.Name)
		}
	}
	return
}

func (prv *Provider) Get(port string) (sio monitor.MonitorIO, err error) {
	sio, ok := prv.Ports[port]
	if !ok {
		err = fmt.Errorf("%s not found", port)
	}
	return
}

func (prv *Provider) Update(filters ...string) (list []string) {
	if prv.Ports == nil {
		prv.Ports = make(map[string]*SerialIO)
	}
	if len(filters) > 0 {
		prv.Filter = filters
	} else {
		prv.Filter = defaultFilters
	}

	list = prv.ListPorts()
	inList := func(s string) bool {
		for _, l := range list {
			if s == l {
				return true
			}
		}
		return false
	}

	for _, l := range list {
		_, ok := prv.Ports[l]
		if !ok {
			prv.Ports[l] = &SerialIO{
				Name:      l,
				available: true,
			}
		}
	}

	for c, p := range prv.Ports {
		if !inList(c) {
			p.available = false
		} else if !p.available {
			p.available = true
			p.opened = false
		}
	}

	list = prv.List()
	return
}

func (prv *Provider) SetFilter(newFilter []string) {
	prv.Filter = newFilter
}

func (prv *Provider) includedPort(fname string) bool {
	for _, f := range prv.Filter {
		if strings.Contains(fname, f) {
			return true
		}
	}
	return false
}

func (prv *Provider) ListPorts() (ports []string) {
	ports = make([]string, 0)
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
		if prv.includedPort(name) {
			ports = append(ports, dev+name)
		}
	}
	return
}
