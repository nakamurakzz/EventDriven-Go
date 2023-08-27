package main

import (
	"time"

	"github.com/nakamurakzz/event-driven-go/hub"
	"github.com/nakamurakzz/event-driven-go/sendor"
	"github.com/nakamurakzz/event-driven-go/sensor"
)

func main() {
	envSendor := sendor.NewEnvSendor()
	hub := hub.NewHub()
	envSensor := sensor.NewEnvSensorer()

	hub.Register(&envSendor)
	envSensor.Register(&hub)

	// 繰り返し実行
	for {
		envSensor.SetTemplature(10)
		envSensor.Notify()
		time.Sleep(5 * time.Second)
	}
}
