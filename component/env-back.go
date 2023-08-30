package component

import (
	"log"
	"net/http"
	"strconv"
)

var _ Componenter = (*EnvBackComponent)(nil)

func NewEnvBackComponent() Componenter {
	return &EnvBackComponent{
		Type: []int{EnvEventFromFront},
	}
}

func (e *EnvBackComponent) Register(h *Huber) {
	e.hubers = append(e.hubers, h)
}

func (e *EnvBackComponent) Recieve(data interface{}) {
	// 型アサーション
	sData, ok := data.(EnvPayload)
	if !ok {
		log.Printf("Failed to type assert data: %v to EnvPayload", data)
		return
	}
	e.temperature = sData.Temperature
	e.Print()
}

func (e *EnvBackComponent) Print() {
	log.Printf("Temperature: %f", e.temperature)
}

func (e *EnvBackComponent) GetType() []int {
	return e.Type
}

func (e *EnvBackComponent) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/env", e.GetFromHttp)
	log.Println("Start env EnvBackComponent port:7002")
	return http.ListenAndServe(":7002", mux)
}

func (e *EnvBackComponent) GetFromHttp(w http.ResponseWriter, r *http.Request) {
	temperature, err := strconv.ParseFloat(r.URL.Query().Get("t"), 64)
	if err != nil {
		http.Error(w, "Invalid env value", http.StatusBadRequest)
		return
	}
	e.SetTemperature(temperature)
	e.Notify()
}

func (e *EnvBackComponent) Notify() {
	payload := NewReceivePayload(EnvEventFromBack, EnvPayload{
		Temperature: e.temperature,
	})

	for _, ob := range e.hubers {
		go (*ob).Receive(payload)
	}
}

func (e *EnvBackComponent) SetTemperature(t float64) {
	e.temperature = t
}
