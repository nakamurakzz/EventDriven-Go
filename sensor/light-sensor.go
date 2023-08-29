package sensor

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nakamurakzz/event-driven-go/hub"
	"github.com/nakamurakzz/event-driven-go/types"
)

type LightSensorer struct {
	power  float64
	Type   string
	hubers []*hub.Huber
}

var _ Sensorer = (*LightSensorer)(nil)

func NewLightSensor() *LightSensorer {
	return &LightSensorer{
		Type: "light",
	}
}

func (l *LightSensorer) Register(h *hub.Huber) {
	l.hubers = append(l.hubers, h)
}

func (l *LightSensorer) Notify() {
	payload := hub.NewReceivePayload(l.Type, types.LightSensorPayload{
		Power: l.power,
	})

	for _, ob := range l.hubers {
		go (*ob).Receive(payload)
	}
}

func (l *LightSensorer) GetPower() float64 {
	return l.power
}

func (l *LightSensorer) SetPower(p float64) {
	l.power = p
}

func (l *LightSensorer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/light", l.GetFromHttp)
	log.Println("Start light sensor")
	return http.ListenAndServe(":6001", mux)
}

func (l *LightSensorer) GetFromHttp(w http.ResponseWriter, r *http.Request) {
	power, err := strconv.ParseFloat(r.URL.Query().Get("p"), 64)
	if err != nil {
		http.Error(w, "Invalid power value", http.StatusBadRequest)
		return
	}
	l.SetPower(power)
	l.Notify()
}

func (l *LightSensorer) Recieve(data interface{}) {
	// 型アサーション
	sData, ok := data.(types.LightSensorPayload)
	if !ok {
		log.Printf("Failed to type assert data: %v to LightSendorPayload", data)
		return
	}
	l.power = sData.Power
	log.Printf("Power: %f", l.power)
}

func (l *LightSensorer) GetType() string {
	return l.Type
}
