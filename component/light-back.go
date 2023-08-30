package component

import (
	"log"
	"net/http"
	"strconv"
)

var _ Componenter = (*LightBackComponent)(nil)

func NewLightBackComponent() *LightBackComponent {
	return &LightBackComponent{
		Type: []int{LightEventFromFront},
	}
}

func (l *LightBackComponent) Register(h *Huber) {
	l.hubers = append(l.hubers, h)
}

func (l *LightBackComponent) Recieve(data interface{}) {
	// 型アサーション
	sData, ok := data.(LightPayload)
	if !ok {
		log.Printf("Failed to type assert data: %v to LightPayload", data)
		return
	}
	l.power = sData.Power
	l.Print()
}

func (l *LightBackComponent) Print() {
	log.Printf("Power: %f", l.power)
}

func (l *LightBackComponent) GetType() []int {
	return l.Type
}

func (l *LightBackComponent) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/light", l.getFromHttp)
	log.Println("Start light LightBackComponent port:7001")
	return http.ListenAndServe(":7001", mux)
}

func (l *LightBackComponent) getFromHttp(w http.ResponseWriter, r *http.Request) {
	power, err := strconv.ParseFloat(r.URL.Query().Get("p"), 64)
	if err != nil {
		http.Error(w, "Invalid power value", http.StatusBadRequest)
		return
	}
	l.SetPower(power)
	l.Notify()
}

func (l *LightBackComponent) Notify() {
	payload := NewReceivePayload(LightEventFromBack, LightPayload{
		Power: l.power,
	})

	for _, ob := range l.hubers {
		go (*ob).Receive(payload)
	}
}

func (l *LightBackComponent) SetPower(p float64) {
	l.power = p
}
