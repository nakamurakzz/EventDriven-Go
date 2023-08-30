package component

import (
	"log"
	"net/http"
	"strconv"
)

// assert type
var _ Componenter = (*EnvFrontComponent)(nil)

func NewEnvFrontComponent() *EnvFrontComponent {
	return &EnvFrontComponent{
		Type: []int{EnvEventFromBack},
	}
}

func (e *EnvFrontComponent) Register(h *Huber) {
	e.hubers = append(e.hubers, h)
}

func (e *EnvFrontComponent) Notify() {
	payload := NewReceivePayload(EnvEventFromFront, EnvPayload{
		Temperature: e.temperature,
	})

	for _, ob := range e.hubers {
		go (*ob).Receive(payload)
	}
}

func (e *EnvFrontComponent) SetTemperature(t float64) {
	e.temperature = t
}

func (e *EnvFrontComponent) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/env", e.getFromHttp)
	log.Println("Start env frontComponent port:6002")
	return http.ListenAndServe(":6002", mux)
}

func (e *EnvFrontComponent) getFromHttp(w http.ResponseWriter, r *http.Request) {
	temperature, err := strconv.ParseFloat(r.URL.Query().Get("t"), 64)
	if err != nil {
		http.Error(w, "Invalid temperature value", http.StatusBadRequest)
		return
	}
	e.SetTemperature(temperature)
	e.Notify()
}

func (l *EnvFrontComponent) Recieve(data interface{}) {
	// 型アサーション
	sData, ok := data.(EnvPayload)
	if !ok {
		log.Printf("Failed to type assert data: %v to EnvPayload", data)
		return
	}
	l.temperature = sData.Temperature
	log.Printf("temperature: %f", l.temperature)
}

func (e *EnvFrontComponent) GetType() []int {
	return e.Type
}
