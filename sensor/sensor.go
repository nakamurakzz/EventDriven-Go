package sensor

import (
	"net/http"

	"github.com/nakamurakzz/event-driven-go/hub"
)

type Sensorer interface {
	Register(h *hub.Huber)
	Notify()
	Recieve(data interface{})

	Start() error
	GetFromHttp(w http.ResponseWriter, r *http.Request)
	GetType() string
}
