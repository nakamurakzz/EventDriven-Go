package sensor

import (
	"net/http"

	"github.com/nakamurakzz/event-driven-go/hub"
)

type Sensorer interface {
	hub.Observer

	Start() error
	GetFromHttp(w http.ResponseWriter, r *http.Request)
}
