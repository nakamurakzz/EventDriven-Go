package observer

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nakamurakzz/event-driven-go/hub"
	"github.com/nakamurakzz/event-driven-go/types"
)

type LightFrontObserver struct {
	power  float64
	Type   []int
	hubers []*hub.Huber
}

var _ Observer = (*LightFrontObserver)(nil)

func NewLightFrontObserver() *LightFrontObserver {
	return &LightFrontObserver{
		Type: []int{hub.LightEventFromBack},
	}
}

func (l *LightFrontObserver) Register(h *hub.Huber) {
	l.hubers = append(l.hubers, h)
}

func (l *LightFrontObserver) Notify() {
	payload := hub.NewReceivePayload(hub.LightEventFromFront, types.LightPayload{
		Power: l.power,
	})

	for _, ob := range l.hubers {
		go (*ob).Receive(payload)
	}
}

func (l *LightFrontObserver) GetPower() float64 {
	return l.power
}

func (l *LightFrontObserver) SetPower(p float64) {
	l.power = p
}

func (l *LightFrontObserver) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/light", l.GetFromHttp)
	log.Println("Start light frontObserver port:6001")
	return http.ListenAndServe(":6001", mux)
}

func (l *LightFrontObserver) GetFromHttp(w http.ResponseWriter, r *http.Request) {
	power, err := strconv.ParseFloat(r.URL.Query().Get("p"), 64)
	if err != nil {
		http.Error(w, "Invalid power value", http.StatusBadRequest)
		return
	}
	l.SetPower(power)
	l.Notify()
}

func (l *LightFrontObserver) Recieve(data interface{}) {
	// 型アサーション
	sData, ok := data.(types.LightPayload)
	if !ok {
		log.Printf("Failed to type assert data: %v to LightPayload", data)
		return
	}
	l.power = sData.Power
	log.Printf("Power: %f", l.power)
}

func (l *LightFrontObserver) GetType() []int {
	return l.Type
}
