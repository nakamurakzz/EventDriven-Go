package sensor

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nakamurakzz/event-driven-go/hub"
)

type LightSensorer struct {
	power      float64
	sensorType string

	hubers []*hub.Huber
}

func NewLightSensorer() *LightSensorer {
	return &LightSensorer{
		sensorType: "light",
	}
}

func (l *LightSensorer) Register(h *hub.Huber) {
	l.hubers = append(l.hubers, h)
}

func (l *LightSensorer) Notify() {
	log.Println("Notify")
	payload := hub.NewRecievePayload(l.sensorType, hub.LightSensorPayload{
		Power: l.power,
	})

	log.Printf("payload: %v", payload)

	for _, ob := range l.hubers {
		go (*ob).Recieve(payload)
	}
}

func (l *LightSensorer) GetTemplature() float64 {
	log.Println("GetTemplature")
	return l.power
}

func (l *LightSensorer) SetTemplature(t float64) {
	log.Printf("SetTemplature: %f", t)
	l.power = t
}

func (l *LightSensorer) Start() error {
	// HTTPサーバーを起動
	mux := http.NewServeMux()
	mux.HandleFunc("/light", l.Recieve)
	// Web サーバーの待ち受けを開始
	log.Fatal(http.ListenAndServe(":6001", mux))
	return nil
}

func (l *LightSensorer) Recieve(w http.ResponseWriter, r *http.Request) {
	log.Println("Recieve")
	power, err := strconv.ParseFloat(r.URL.Query().Get("p"), 64)
	if err != nil {
		log.Fatal(err)
	}
	l.power = power
	l.Notify()
}
