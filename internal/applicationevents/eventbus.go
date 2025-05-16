package applicationevents

import (
	"sync"
)

type EventBus struct {
	handlers map[string][]ApplicationEventHandler
	mu       sync.RWMutex
}

// NOTE: ideally event bus should be a singleton
func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]ApplicationEventHandler),
	}
}

func (b *EventBus) Subscribe(event string, handler ApplicationEventHandler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[event] = append(b.handlers[event], handler)
}

// NOTE:  ideally event bus should should have a queue, for now I will directly call the handlers
func (b *EventBus) Publish(event string, payload any) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, handler := range b.handlers[event] {
		go handler.HandleEvent(event, payload)
	}
}
