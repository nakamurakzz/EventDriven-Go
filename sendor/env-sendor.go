package sendor

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nakamurakzz/event-driven-go/hub"
	"github.com/nakamurakzz/event-driven-go/types"
)

type EnvSendor struct {
	Type        string
	temperature float64
	hubers      []*hub.Huber
}

// 型アサーション
var _ Sendorer = (*EnvSendor)(nil)

func NewEnvSendor() Sendorer {
	return &EnvSendor{
		Type: "env",
	}
}

func (e *EnvSendor) Register(h *hub.Huber) {
	e.hubers = append(e.hubers, h)
}

func (e *EnvSendor) Recieve(data interface{}) {
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

func (e *EnvSendor) GetType() string {
	return e.Type
}

func (e *EnvSendor) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/light", e.GetFromHttp)
	log.Println("Start light sensor")
	return http.ListenAndServe(":7002", mux)
}

func (e *EnvSendor) GetFromHttp(w http.ResponseWriter, r *http.Request) {
	power, err := strconv.ParseFloat(r.URL.Query().Get("p"), 64)
	if err != nil {
		http.Error(w, "Invalid power value", http.StatusBadRequest)
		return
	}
	e.SetPower(power)
	e.Notify()
}

func (e *EnvSendor) Notify() {
	payload := hub.NewReceivePayload(e.Type, types.LightSensorPayload{
		Power: e.temperature,
	})

	for _, ob := range e.hubers {
		go (*ob).Receive(payload)
	}
}

func (e *EnvSendor) SetPower(t float64) {
	e.temperature = t
}
