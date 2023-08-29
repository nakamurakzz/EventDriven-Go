package sendor

import (
	"log"

	"github.com/nakamurakzz/event-driven-go/types"
)

type LightSendor struct {
	sendorType string
	power      float64
}

// 型アサーション
var _ Sendorer = (*LightSendor)(nil)

func NewLightSendor() Sendorer {
	return &LightSendor{
		sendorType: "light",
	}
}

func (l *LightSendor) Receive(data interface{}) {
	// 型アサーション
	sData, ok := data.(types.LightSensorPayload)
	if !ok {
		log.Printf("Failed to type assert data: %v to LightSendorPayload", data)
		return
	}
	l.power = sData.Power
	l.Print()
}

func (l *LightSendor) Print() {
	log.Printf("Power: %f", l.power)
}

func (l *LightSendor) GetType() string {
	return l.sendorType
}
