package sensor

import (
	"github.com/nakamurakzz/event-driven-go/hub"
)

type Sensorer interface {
	Resister(hub hub.Huber)
	Notify()

	GetTemplature() float64
}
