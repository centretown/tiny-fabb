package monitor

type Broker interface {
	Start()
	Ask() (port string, err error)
}
