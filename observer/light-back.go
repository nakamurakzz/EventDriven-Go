package observer

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nakamurakzz/event-driven-go/hub"
	"github.com/nakamurakzz/event-driven-go/types"
)

type LightBackObserver struct {
	power  float64
	Type   []int
	hubers []*hub.Huber
}

// 型アサーション
var _ Observer = (*LightBackObserver)(nil)

func NewLightBackObserver() Observer {
	return &LightBackObserver{
		Type: []int{hub.LightEventFromFront},
	}
}

func (l *LightBackObserver) Register(h *hub.Huber) {
	l.hubers = append(l.hubers, h)
}

func (l *LightBackObserver) Recieve(data interface{}) {
	// 型アサーション
	sData, ok := data.(types.LightPayload)
	if !ok {
		log.Printf("Failed to type assert data: %v to LightPayload", data)
		return
	}
	l.power = sData.Power
	l.Print()
}

func (l *LightBackObserver) Print() {
	log.Printf("Power: %f", l.power)
}

func (l *LightBackObserver) GetType() []int {
	return l.Type
}

func (l *LightBackObserver) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/light", l.GetFromHttp)
	log.Println("Start light LightBackObserver port:7001")
	return http.ListenAndServe(":7001", mux)
}

func (l *LightBackObserver) GetFromHttp(w http.ResponseWriter, r *http.Request) {
	power, err := strconv.ParseFloat(r.URL.Query().Get("p"), 64)
	if err != nil {
		http.Error(w, "Invalid power value", http.StatusBadRequest)
		return
	}
	l.SetPower(power)
	l.Notify()
}

func (l *LightBackObserver) Notify() {
	payload := hub.NewReceivePayload(hub.LightEventFromBack, types.LightPayload{
		Power: l.power,
	})

	for _, ob := range l.hubers {
		go (*ob).Receive(payload)
	}
}

func (l *LightBackObserver) SetPower(p float64) {
	l.power = p
}
