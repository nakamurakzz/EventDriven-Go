package hub

import "fmt"

// EventTypes
const (
	LightEventFromBack = iota
	EnvEventFromBack
	LightEventFromFront
	EnvEventFromFront
	LightEvent
	EnvEvent
)

type Observer interface {
	Register(h *Huber)
	Notify()
	Recieve(data interface{})
	GetType() []int
	Start() error
}

type ReceivePayload struct {
	eventType int
	data      interface{}
}

func NewReceivePayload(eventType int, data interface{}) ReceivePayload {
	return ReceivePayload{
		eventType: eventType,
		data:      data,
	}
}

type Huber interface {
	Notify(payload ReceivePayload)
	Receive(payload ReceivePayload)
	Register(o Observer)
}

type Hub struct {
	observers map[int][]Observer
}

func NewHub() Huber {
	return &Hub{
		observers: make(map[int][]Observer),
	}
}

func (h *Hub) Register(o Observer) {
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
