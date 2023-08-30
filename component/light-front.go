package component

import (
	"log"
	"net/http"
	"strconv"
)

var _ Componenter = (*LightFrontComponent)(nil)

func NewLightFrontComponent() *LightFrontComponent {
	return &LightFrontComponent{
		Type: []int{LightEventFromBack},
	}
}

func (l *LightFrontComponent) Register(h *Huber) {
	l.hubers = append(l.hubers, h)
}

func (l *LightFrontComponent) Notify() {
	payload := NewReceivePayload(LightEventFromFront, LightPayload{
		Power: l.power,
	})

	for _, ob := range l.hubers {
		go (*ob).Receive(payload)
	}
}

func (l *LightFrontComponent) GetPower() float64 {
	return l.power
}

func (l *LightFrontComponent) SetPower(p float64) {
	l.power = p
}

func (l *LightFrontComponent) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/light", l.getFromHttp)
	log.Println("Start light frontComponent port:6001")
	return http.ListenAndServe(":6001", mux)
}

func (l *LightFrontComponent) getFromHttp(w http.ResponseWriter, r *http.Request) {
	power, err := strconv.ParseFloat(r.URL.Query().Get("p"), 64)
	if err != nil {
		http.Error(w, "Invalid power value", http.StatusBadRequest)
		return
	}
	l.SetPower(power)
	l.Notify()
}

func (l *LightFrontComponent) Recieve(data interface{}) {
	// 型アサーション
	sData, ok := data.(LightPayload)
	if !ok {
		log.Printf("Failed to type assert data: %v to LightPayload", data)
		return
	}
	l.power = sData.Power
	log.Printf("Power: %f", l.power)
}

func (l *LightFrontComponent) GetType() []int {
	return l.Type
}
