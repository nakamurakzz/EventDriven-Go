package component

import (
	"fmt"
)

// EventTypes
const (
	LightEventFromBack = iota
	EnvEventFromBack
	LightEventFromFront
	EnvEventFromFront
	LightEvent
	EnvEvent
)

func NewReceivePayload(eventType int, data interface{}) ReceivePayload {
	return ReceivePayload{
		eventType: eventType,
		data:      data,
	}
}

func NewHub() Huber {
	return &Hub{
		observers: make(map[int][]Componenter),
	}
}

var _ Huber = (*Hub)(nil)

func (h *Hub) Register(o Componenter) {
	for _, t := range o.GetType() {
		h.observers[t] = append(h.observers[t], o)
	}
}

func (h *Hub) Receive(payload ReceivePayload) {
	h.Notify(payload)
}

func (h *Hub) Notify(payload ReceivePayload) {
	fmt.Printf("Notify: %v\n", payload)
	for _, o := range h.observers[payload.eventType] {
		go o.Recieve(payload.data)
	}
}
