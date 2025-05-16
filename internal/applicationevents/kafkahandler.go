package applicationevents

import "fmt"

type KafkaApplicationEventHandler struct {
	EventBus *EventBus
}

func NewKafkaApplicationEventHandler(eventBus *EventBus) EventHandler {
	return &KafkaApplicationEventHandler{EventBus: eventBus}
}

func (h *KafkaApplicationEventHandler) HandleEvent(event string, payload any) {
	fmt.Printf("Received event: %s, payload: %v\n", event, payload)
	//TODO: in real app, we should publish the event to the kafka topic
}
