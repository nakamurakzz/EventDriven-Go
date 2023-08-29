package sensor

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nakamurakzz/event-driven-go/hub"
	"github.com/nakamurakzz/event-driven-go/types"
)

type EnvSensorer struct {
	temperature float64
	Type        string
	hubers      []*hub.Huber
}

// assert type
var _ Sensorer = (*EnvSensorer)(nil)

func NewEnvSensor() *EnvSensorer {
	return &EnvSensorer{
		Type: "env",
	}
}

func (e *EnvSensorer) Register(h *hub.Huber) {
	e.hubers = append(e.hubers, h)
}

func (e *EnvSensorer) Notify() {
	payload := hub.NewReceivePayload(e.Type, types.EnvSensorPayload{
		Temperature: e.temperature,
	})

	for _, ob := range e.hubers {
		go (*ob).Receive(payload)
	}
}

func (e *EnvSensorer) SetTemperature(t float64) {
	e.temperature = t
}

func (e *EnvSensorer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/env", e.GetFromHttp)
	log.Println("Start env sensor")
	return http.ListenAndServe(":6002", mux)
}

func (e *EnvSensorer) GetFromHttp(w http.ResponseWriter, r *http.Request) {
	temperature, err := strconv.ParseFloat(r.URL.Query().Get("t"), 64)
	if err != nil {
		http.Error(w, "Invalid temperature value", http.StatusBadRequest)
		return
	}
	e.SetTemperature(temperature)
	e.Notify()
}

func (l *EnvSensorer) Recieve(data interface{}) {
	// 型アサーション
	sData, ok := data.(types.LightSensorPayload)
	if !ok {
		log.Printf("Failed to type assert data: %v to LightSendorPayload", data)
		return
	}
	l.temperature = sData.Power
	log.Printf("temperature: %f", l.temperature)
}

func (e *EnvSensorer) GetType() string {
	return e.Type
}
