package hub

import (
	"github.com/nakamurakzz/event-driven-go/sendor"
)

type ReceivePayload struct {
	eventType string
	data      interface{}
}

func NewReceivePayload(eventType string, data interface{}) ReceivePayload {
	return ReceivePayload{
		eventType: eventType,
		data:      data,
	}
}

type Huber interface {
	Notify(payload ReceivePayload)
	Receive(payload ReceivePayload)
	Register(s sendor.Sendorer)
}

type Hub struct {
	sendors map[string]sendor.Sendorer
}

func NewHub() Huber {
	return &Hub{
		sendors: make(map[string]sendor.Sendorer),
	}
}

func (h *Hub) Register(s sendor.Sendorer) {
	h.sendors[s.GetSendorType()] = s
}

func (h *Hub) Receive(payload ReceivePayload) {
	h.Notify(payload)
}

func (h *Hub) Notify(payload ReceivePayload) {
	if s, ok := h.sendors[payload.eventType]; ok {
		s.Receive(payload.data)
	}
}
