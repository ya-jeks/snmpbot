package server

import (
	"fmt"
	"github.com/qmsk/snmpbot/client"
	"sync"
)

func newEngine(clientEngine *client.Engine) *Engine {
	return &Engine{
		clientEngine: clientEngine,
		mibs:         AllMIBs(),
		hosts:        make(Hosts),
	}
}

type Engine struct {
	clientEngine  *client.Engine
	clientOptions client.Options

	mibs       MIBs
	hosts      Hosts
	hostsMutex sync.Mutex
}

func (engine *Engine) loadConfig(config Config) error {
	engine.clientOptions = config.ClientOptions

	for hostName, hostConfig := range config.Hosts {
		if host, err := newHost(engine, HostID(hostName), hostConfig); err != nil {
			return fmt.Errorf("Failed to load host %v: %v", hostName, err)
		} else {
			host.start()

			engine.hosts[host.id] = host
		}
	}

	return nil
}

func (engine *Engine) MIBs() MIBs {
	return engine.mibs
}

func (engine *Engine) AddHost(host *Host) {
	engine.hostsMutex.Lock()
	defer engine.hostsMutex.Unlock()

	engine.hosts[host.id] = host
}

func (engine *Engine) DelHost(host *Host) bool {
	engine.hostsMutex.Lock()
	defer engine.hostsMutex.Unlock()

	if _, exists := engine.hosts[host.id]; exists {
		delete(engine.hosts, host.id)
		return true
	} else {
		return false
	}
}

func (engine *Engine) Hosts() Hosts {
	var hosts = make(Hosts)

	engine.hostsMutex.Lock()
	defer engine.hostsMutex.Unlock()

	for hostID, host := range engine.hosts {
		if !host.IsUp() {
			continue
		}
		hosts[hostID] = host
	}

	return hosts
}

func (engine *Engine) Objects() Objects {
	// TODO: limit by MIBs?
	return AllObjects()
}

func (engine *Engine) Tables() Tables {
	// TODO: limit by MIBs?
	return AllTables()
}

func (engine *Engine) QueryObjects(q ObjectQuery) <-chan ObjectResult {
	q.resultChan = make(chan ObjectResult)

	go q.query()

	return q.resultChan
}

func (engine *Engine) QueryTables(q TableQuery) <-chan TableResult {
	q.resultChan = make(chan TableResult)

	go q.query()

	return q.resultChan
}
