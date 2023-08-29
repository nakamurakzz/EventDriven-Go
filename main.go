package main

import (
	"github.com/nakamurakzz/event-driven-go/hub"
	"github.com/nakamurakzz/event-driven-go/sendor"
	"github.com/nakamurakzz/event-driven-go/sensor"
	"golang.org/x/sync/errgroup"
)

func main() {
	h := hub.NewHub()

	observers := initializeObserver()
	for _, o := range observers {
		o.Register(&h)
		h.Register(o)
	}

	g := errgroup.Group{}
	for _, observer := range observers {
		o := observer // capture loop variable
		g.Go(func() error {
			return o.Start()
		})
	}

	g.Wait()
}

func initializeObserver() []hub.Observer {
	return []hub.Observer{
		sendor.NewEnvSendor(),
		sendor.NewLightSendor(),
		sensor.NewEnvSensor(),
		sensor.NewLightSensor(),
	}
}
