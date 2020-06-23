package subscribers

import (
	"github.com/epc-blockchain/go-xpx-chain-sdk/sdk"
	"sync"
)

type (
	StatusHandler func(*sdk.StatusInfo) bool

	Status interface {
		AddHandlers(address *sdk.Address, handlers ...StatusHandler) error
		RemoveHandlers(address *sdk.Address, handlers ...*StatusHandler) bool
		HasHandlers(address *sdk.Address) bool
		GetHandlers(address *sdk.Address) []*StatusHandler
		GetAddresses() []string
	}

	statusImpl struct {
		sync.Mutex
		newSubscriberCh    chan *statusSubscription
		removeSubscriberCh chan *statusSubscription
		subscribers        map[string][]*StatusHandler
	}
	statusSubscription struct {
		address  *sdk.Address
		handlers []*StatusHandler
		resultCh chan bool
	}
)

func NewStatus() Status {

	p := &statusImpl{
		subscribers:        make(map[string][]*StatusHandler),
		newSubscriberCh:    make(chan *statusSubscription),
		removeSubscriberCh: make(chan *statusSubscription),
	}
	go p.handleNewSubscription()
	return p
}

func (e *statusImpl) handleNewSubscription() {
	for {
		select {
		case s := <-e.newSubscriberCh:
			e.addSubscription(s)
		case s := <-e.removeSubscriberCh:
			e.removeSubscription(s)
		}
	}
}

func (e *statusImpl) addSubscription(s *statusSubscription) {
	e.Lock()
	defer e.Unlock()
	if _, ok := e.subscribers[s.address.Address]; !ok {
		e.subscribers[s.address.Address] = make([]*StatusHandler, 0)
	}
	for i := 0; i < len(s.handlers); i++ {
		e.subscribers[s.address.Address] = append(e.subscribers[s.address.Address], s.handlers[i])
	}
}

func (e *statusImpl) removeSubscription(s *statusSubscription) {
	e.Lock()
	defer e.Unlock()
	if external, ok := e.subscribers[s.address.Address]; !ok || len(external) == 0 {
		s.resultCh <- false
	}

	itemCount := len(e.subscribers[s.address.Address])
	for _, removeHandler := range s.handlers {
		for index, currentHandlers := range e.subscribers[s.address.Address] {
			if removeHandler == currentHandlers {
				e.subscribers[s.address.Address] = append(e.subscribers[s.address.Address][:index],
					e.subscribers[s.address.Address][index+1:]...)
			}
		}
	}

	s.resultCh <- itemCount != len(e.subscribers[s.address.Address])
}

func (e *statusImpl) AddHandlers(address *sdk.Address, handlers ...StatusHandler) error {

	if len(handlers) == 0 {
		return nil
	}

	refHandlers := make([]*StatusHandler, len(handlers))
	for i, h := range handlers {
		refHandlers[i] = &h
	}

	e.newSubscriberCh <- &statusSubscription{
		address:  address,
		handlers: refHandlers,
	}
	return nil
}

func (e *statusImpl) RemoveHandlers(address *sdk.Address, handlers ...*StatusHandler) bool {
	if len(handlers) == 0 {
		return false
	}

	resCh := make(chan bool)
	e.removeSubscriberCh <- &statusSubscription{
		address:  address,
		handlers: handlers,
		resultCh: resCh,
	}

	return <-resCh
}

func (e *statusImpl) HasHandlers(address *sdk.Address) bool {
	e.Lock()
	defer e.Unlock()
	return len(e.subscribers[address.Address]) > 0 && e.subscribers[address.Address] != nil
}

func (e *statusImpl) GetHandlers(address *sdk.Address) []*StatusHandler {
	e.Lock()
	defer e.Unlock()
	if res, ok := e.subscribers[address.Address]; ok && res != nil {
		return res
	}

	return nil
}

func (e *statusImpl) GetAddresses() []string {
	e.Lock()
	defer e.Unlock()
	addresses := make([]string, 0, len(e.subscribers))
	for addr := range e.subscribers {
		addresses = append(addresses, addr)
	}

	return addresses
}
