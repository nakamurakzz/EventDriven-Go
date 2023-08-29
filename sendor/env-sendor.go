package sendor

import (
	"log"

	"github.com/nakamurakzz/event-driven-go/types"
)

type EnvSendor struct {
	sendorType  string
	temperature float64
}

// 型アサーション
var _ Sendorer = (*EnvSendor)(nil)

func NewEnvSendor() Sendorer {
	return &EnvSendor{
		sendorType: "env",
	}
}

func (e *EnvSendor) Receive(data interface{}) {
	log.Printf("Receive: %v", data)
	// 型アサーション
	sData, ok := data.(types.EnvSensorPayload)
	if !ok {
		log.Printf("Failed to type assert data: %v to EnvSensorPayload", data)
		return
	}
	e.temperature = sData.Temperature
	e.Print()
}

func (e *EnvSendor) Print() {
	log.Printf("Temperature: %f", e.temperature)
}

func (e *EnvSendor) GetSendorType() string {
	return e.sendorType
}
