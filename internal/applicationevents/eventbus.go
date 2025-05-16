package applicationevents

import (
	"log"
	"sync"
)

type EventBus struct {
	handlers map[ApplicationEvent][]ApplicationEventHandler
	mu       sync.RWMutex
}

// NOTE: ideally event bus should be a singleton
func NewEventBus() *EventBus {
	log.Println("Initializing EventBus")
	return &EventBus{
		handlers: make(map[ApplicationEvent][]ApplicationEventHandler),
	}
}

func (b *EventBus) Subscribe(event ApplicationEvent, handler ApplicationEventHandler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[event] = append(b.handlers[event], handler)
	log.Printf("Subscribed to event: %s", event)
}

// NOTE:  ideally event bus should should have a queue, for now I will directly call the handlers
func (b *EventBus) Publish(event ApplicationEvent, payload any) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, handler := range b.handlers[event] {
		go handler.HandleEvent(event, payload)
	}
}
