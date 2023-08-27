package sendor

import (
	"log"
)

type LightSendor struct {
	sendorType string
	power      float64
}

// 型アサーション
var _ Sendor = (*LightSendor)(nil)

func NewLightSendor() Sendor {
	return &LightSendor{
		sendorType: "light",
	}
}

type LightSendorPayload struct {
	Power float64
}

func (e *LightSendor) Recieve(data interface{}) {
	log.Printf("Recieve: %v", data)

	// 型アサーション
	sData, ok := data.(LightSendorPayload)
	if !ok {
		log.Printf("data: %v", data)
		log.Println("failed to type assertion")
		return
	}
	e.power = sData.Power
	e.Print()
}

func (e *LightSendor) Print() {
	log.Println("power: ", e.power)
}

func (e *LightSendor) GetSendorType() string {
	log.Println("GetSendorType")
	return e.sendorType
}
