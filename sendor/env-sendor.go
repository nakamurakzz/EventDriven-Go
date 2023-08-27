package sendor

import (
	"log"
)

type EnvSendor struct {
	sendorType  string
	temperature float64
}

// 型アサーション
var _ Sendor = (*EnvSendor)(nil)

func NewEnvSendor() Sendor {
	return &EnvSendor{
		sendorType: "env",
	}
}

type EnvSendorPayload struct {
	Temperature float64
}

func (e *EnvSendor) Recieve(data interface{}) {
	log.Printf("Recieve: %v", data)

	// 型アサーション
	sData, ok := data.(EnvSendorPayload)
	if !ok {
		log.Printf("data: %v", data)
		log.Println("failed to type assertion")
		return
	}
	e.temperature = sData.Temperature
	e.Print()
}

func (e *EnvSendor) Print() {
	log.Println("temperature: ", e.temperature)
}

func (e *EnvSendor) GetSendorType() string {
	log.Println("GetSendorType")
	return e.sendorType
}
