package observer

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nakamurakzz/event-driven-go/hub"
	"github.com/nakamurakzz/event-driven-go/types"
)

type EnvFrontObserver struct {
	temperature float64
	Type        []int
	hubers      []*hub.Huber
}

// assert type
var _ Observer = (*EnvFrontObserver)(nil)

func NewEnvFrontObserver() *EnvFrontObserver {
	return &EnvFrontObserver{
		Type: []int{hub.EnvEventFromBack},
	}
}

func (e *EnvFrontObserver) Register(h *hub.Huber) {
	e.hubers = append(e.hubers, h)
}

func (e *EnvFrontObserver) Notify() {
	payload := hub.NewReceivePayload(hub.EnvEventFromFront, types.EnvPayload{
		Temperature: e.temperature,
	})

	for _, ob := range e.hubers {
		go (*ob).Receive(payload)
	}
}

func (e *EnvFrontObserver) SetTemperature(t float64) {
	e.temperature = t
}

func (e *EnvFrontObserver) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/env", e.GetFromHttp)
	log.Println("Start env frontObserver port:6002")
	return http.ListenAndServe(":6002", mux)
}

func (e *EnvFrontObserver) GetFromHttp(w http.ResponseWriter, r *http.Request) {
	temperature, err := strconv.ParseFloat(r.URL.Query().Get("t"), 64)
	if err != nil {
		http.Error(w, "Invalid temperature value", http.StatusBadRequest)
		return
	}
	e.SetTemperature(temperature)
	e.Notify()
}

func (l *EnvFrontObserver) Recieve(data interface{}) {
	// 型アサーション
	sData, ok := data.(types.EnvPayload)
	if !ok {
		log.Printf("Failed to type assert data: %v to EnvPayload", data)
		return
	}
	l.temperature = sData.Temperature
	log.Printf("temperature: %f", l.temperature)
}

func (e *EnvFrontObserver) GetType() []int {
	return e.Type
}
