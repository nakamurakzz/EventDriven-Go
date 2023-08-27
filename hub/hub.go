package hub

import (
	"log"

	"github.com/nakamurakzz/event-driven-go/sendor"
)

type RecievePayload struct {
	eventType string
	data      interface{}
}

func NewRecievePayload(eventType string, data interface{}) RecievePayload {
	return RecievePayload{
		eventType: eventType,
		data:      data,
	}
}

type Huber interface {
	Recieve(payload RecievePayload)
	Register(s *sendor.Sendor)
}

type Hub struct {
	sendors []*sendor.Sendor
}

func NewHub() Huber {
	return &Hub{}
}

func (h *Hub) Register(s *sendor.Sendor) {
	log.Printf("Register: %s", (*s).GetSendorType())
	h.sendors = append(h.sendors, s)
}

func (h *Hub) Recieve(payload RecievePayload) {
	log.Printf("Recieve: %v", payload)
	for _, s := range h.sendors {
		if payload.eventType == (*s).GetSendorType() {
			if payload.eventType == "env" {
				sData := sendor.EnvSendorPayload{
					Temperature: payload.data.(EnvSensorPayload).Temperature,
				}
				(*s).Recieve(sData)
				continue
			}
		}
	}
}

type EnvSensorPayload struct {
	Temperature float64
}
