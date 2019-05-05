package handlers

import (
	"github.com/pkg/errors"
	"github.com/proximax-storage/go-xpx-catapult-sdk/sdk"
	"github.com/proximax-storage/go-xpx-catapult-sdk/sdk/websocket/subscribers"
	"sync"
)

func NewStatusHandler(messageMapper sdk.StatusMapper, handlers subscribers.Status) *statusHandler {
	return &statusHandler{
		messageMapper: messageMapper,
		handlers:      handlers,
	}
}

type statusHandler struct {
	messageMapper sdk.StatusMapper
	handlers      subscribers.Status
}

func (h *statusHandler) Handle(address *sdk.Address, resp []byte) bool {
	res, err := h.messageMapper.MapStatus(resp)
	if err != nil {
		panic(errors.Wrap(err, "message mapper error"))
	}

	handlers := h.handlers.GetHandlers(address)
	if len(handlers) == 0 {
		return true
	}

	var wg sync.WaitGroup

	for f := range handlers {
		wg.Add(1)
		go func(f *subscribers.StatusHandler) {
			defer wg.Done()

			callFunc := *f

			if rm := callFunc(res); !rm {
				return
			}

			_, err = h.handlers.RemoveHandlers(address, f)
			if err != nil {
				panic(errors.Wrap(err, "removing handler from storage"))
			}

			return
		}(f)
	}

	wg.Wait()

	return h.handlers.HasHandlers(address)
}