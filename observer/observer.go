package observer

import "github.com/nakamurakzz/event-driven-go/hub"

type Observer interface {
	Register(h *hub.Huber)
	Notify()
	Recieve(data interface{})
	GetType() []int
	Start() error
}
