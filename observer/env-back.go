package observer

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nakamurakzz/event-driven-go/hub"
	"github.com/nakamurakzz/event-driven-go/types"
)

type EnvBackObserver struct {
	temperature float64
	Type        []int
	hubers      []*hub.Huber
}

// 型アサーション
var _ Observer = (*EnvBackObserver)(nil)

func NewEnvBackObserver() Observer {
	return &EnvBackObserver{
		Type: []int{hub.EnvEventFromFront},
	}
}

func (e *EnvBackObserver) Register(h *hub.Huber) {
	e.hubers = append(e.hubers, h)
}

func (e *EnvBackObserver) Recieve(data interface{}) {
	// 型アサーション
	sData, ok := data.(types.EnvPayload)
	if !ok {
		log.Printf("Failed to type assert data: %v to EnvPayload", data)
		return
	}
	e.temperature = sData.Temperature
	e.Print()
}

func (e *EnvBackObserver) Print() {
	log.Printf("Temperature: %f", e.temperature)
}

func (e *EnvBackObserver) GetType() []int {
	return e.Type
}

func (e *EnvBackObserver) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/env", e.GetFromHttp)
	log.Println("Start env EnvBackObserver port:7002")
	return http.ListenAndServe(":7002", mux)
}

func (e *EnvBackObserver) GetFromHttp(w http.ResponseWriter, r *http.Request) {
	temperature, err := strconv.ParseFloat(r.URL.Query().Get("t"), 64)
	if err != nil {
		http.Error(w, "Invalid env value", http.StatusBadRequest)
		return
	}
	e.SetTemperature(temperature)
	e.Notify()
}

func (e *EnvBackObserver) Notify() {
	payload := hub.NewReceivePayload(hub.EnvEventFromBack, types.EnvPayload{
		Temperature: e.temperature,
	})

	for _, ob := range e.hubers {
		go (*ob).Receive(payload)
	}
}

func (e *EnvBackObserver) SetTemperature(t float64) {
	e.temperature = t
}
