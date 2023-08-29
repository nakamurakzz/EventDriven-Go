package main

import (
	"github.com/nakamurakzz/event-driven-go/hub"
	"github.com/nakamurakzz/event-driven-go/sendor"
	"github.com/nakamurakzz/event-driven-go/sensor"
	"golang.org/x/sync/errgroup"
)

func main() {
	h := hub.NewHub()

	sensors := initializeSensors()
	for _, s := range sensors {
		s.Register(&h)
		h.Register(&s)
	}

	sendors := initializeSendors()
	for _, s := range sendors {
		h.Register(s)
	}

	g := errgroup.Group{}
	for _, s := range sensors {
		sensor := s // capture loop variable
		g.Go(func() error {
			return sensor.Start()
		})
	}

	g.Wait()
}

func initializeSensors() []sensor.Sensorer {
	return []sensor.Sensorer{
		sensor.NewEnvSensor(),
		sensor.NewLightSensor(),
	}
}

func initializeSendors() []sendor.Sendorer {
	return []sendor.Sendorer{
		sendor.NewEnvSendor(),
		sendor.NewLightSendor(),
	}
}
