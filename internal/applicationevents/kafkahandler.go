package applicationevents

import (
	"log"
)

type KafkaApplicationEventHandler struct {
	EventBus *EventBus
}

func NewKafkaApplicationEventHandler(eventBus *EventBus) ApplicationEventHandler {
	log.Println("Initializing KafkaApplicationEventHandler")

	for _, event := range ApplicationEvents {
		eventBus.Subscribe(event, &KafkaApplicationEventHandler{EventBus: eventBus})
	}

	return &KafkaApplicationEventHandler{EventBus: eventBus}
}

func (h *KafkaApplicationEventHandler) HandleEvent(event ApplicationEvent, payload any) {
	log.Printf("Received event: %s, payload: %v\n", event, payload)
	//TODO: in real app, we should publish the event to the kafka topic
}
