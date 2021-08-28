package monitor

type Provider interface {
	Update(...string) []string
	List() []string
	Get(string) (MonitorIO, error)
}

type Manager struct {
	providers  []Provider
	connectors []Connector
}

func NewManager() (mgr *Manager) {
	mgr = &Manager{
		providers:  make([]Provider, 0),
		connectors: make([]Connector, 0),
	}
	return
}

func (mgr *Manager) AddProvider(d Provider) {
	mgr.providers = append(mgr.providers, d)
}

func (mgr *Manager) AddConnector(c Connector) {
	mgr.connectors = append(mgr.connectors, c)
}
