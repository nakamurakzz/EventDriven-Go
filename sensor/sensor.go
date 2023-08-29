package sensor

import "github.com/nakamurakzz/event-driven-go/hub"

type Sensorer interface {
	Register(h *hub.Huber)
	Notify()

	Start() error
}
