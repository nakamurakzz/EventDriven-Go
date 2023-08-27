package sensor

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nakamurakzz/event-driven-go/hub"
)

type EnvSensorer struct {
	temperature float64
	sensorType  string

	hubers []*hub.Huber
}

func NewEnvSensorer() *EnvSensorer {
	return &EnvSensorer{
		sensorType: "env",
	}
}

func (e *EnvSensorer) Register(h *hub.Huber) {
	e.hubers = append(e.hubers, h)
}

func (e *EnvSensorer) Notify() {
	log.Println("Notify")
	payload := hub.NewRecievePayload(e.sensorType, hub.EnvSensorPayload{
		Temperature: e.temperature,
	})

	log.Printf("payload: %v", payload)

	for _, ob := range e.hubers {
		go (*ob).Recieve(payload)
	}
}

func (e *EnvSensorer) GetTemplature() float64 {
	log.Println("GetTemplature")
	return e.temperature
}

func (e *EnvSensorer) SetTemplature(t float64) {
	log.Printf("SetTemplature: %f", t)
	e.temperature = t
}

func (e *EnvSensorer) Start() error {
	// HTTPサーバーを起動
	mux := http.NewServeMux()
	mux.HandleFunc("/env", e.Recieve)
	// Web サーバーの待ち受けを開始
	log.Fatal(http.ListenAndServe(":6000", mux))
	return nil
}

func (e *EnvSensorer) Recieve(w http.ResponseWriter, r *http.Request) {
	log.Println("Recieve")
	// get query params and convert to float64
	temperature, err := strconv.ParseFloat(r.URL.Query().Get("t"), 64)
	if err != nil {
		log.Fatal(err)
	}
	e.temperature = temperature
	e.Notify()
}
