package hub

type Observer interface {
	Register(h *Huber)
	Notify()
	Recieve(data interface{})
	GetType() string
}

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
	Register(o Observer)
}

type Hub struct {
	observers map[string][]Observer
}

func NewHub() Huber {
	return &Hub{
		observers: make(map[string][]Observer),
	}
}

func (h *Hub) Register(o Observer) {
	h.observers[o.GetType()] = append(h.observers[o.GetType()], o)
}

func (h *Hub) Receive(payload ReceivePayload) {
	h.Notify(payload)
}

func (h *Hub) Notify(payload ReceivePayload) {
	for _, o := range h.observers[payload.eventType] {
		go o.Recieve(payload.data)
	}
}
