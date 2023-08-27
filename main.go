package main

import (
	"github.com/nakamurakzz/event-driven-go/hub"
	"github.com/nakamurakzz/event-driven-go/sendor"
	"github.com/nakamurakzz/event-driven-go/sensor"
	"golang.org/x/sync/errgroup"
)

func main() {
	hub := hub.NewHub()
	envSendor := sendor.NewEnvSendor()
	envSensor := sensor.NewEnvSensorer()
	lightSendor := sendor.NewLightSendor()
	lightSensor := sensor.NewLightSensorer()

	hub.Register(&envSendor)
	hub.Register(&lightSendor)
	envSensor.Register(&hub)
	lightSensor.Register(&hub)

	errgroup := errgroup.Group{}
	errgroup.Go(func() error {
		return envSensor.Start()
	})
	errgroup.Go(func() error {
		return lightSensor.Start()
	})

	errgroup.Wait()

}
