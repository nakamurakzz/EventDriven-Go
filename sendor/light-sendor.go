package sendor

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nakamurakzz/event-driven-go/hub"
	"github.com/nakamurakzz/event-driven-go/types"
)

type LightSendor struct {
	Type   string
	power  float64
	hubers []*hub.Huber
}

// 型アサーション
var _ Sendorer = (*LightSendor)(nil)

func NewLightSendor() Sendorer {
	return &LightSendor{
		Type: "light",
	}
}

func (l *LightSendor) Register(h *hub.Huber) {
	l.hubers = append(l.hubers, h)
}

func (l *LightSendor) Recieve(data interface{}) {
	// 型アサーション
	sData, ok := data.(types.LightSensorPayload)
	if !ok {
		log.Printf("Failed to type assert data: %v to LightSendorPayload", data)
		return
	}
	l.power = sData.Power
	l.Print()
}

func (l *LightSendor) Print() {
	log.Printf("Power: %f", l.power)
}

func (l *LightSendor) GetType() string {
	return l.Type
}

func (l *LightSendor) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/light", l.GetFromHttp)
	log.Println("Start light sensor")
	return http.ListenAndServe(":7001", mux)
}

func (l *LightSendor) GetFromHttp(w http.ResponseWriter, r *http.Request) {
	power, err := strconv.ParseFloat(r.URL.Query().Get("p"), 64)
	if err != nil {
		http.Error(w, "Invalid power value", http.StatusBadRequest)
		return
	}
	l.SetPower(power)
	l.Notify()
}

func (l *LightSendor) Notify() {
	payload := hub.NewReceivePayload(l.Type, types.LightSensorPayload{
		Power: l.power,
	})

	for _, ob := range l.hubers {
		go (*ob).Receive(payload)
	}
}

func (l *LightSendor) SetPower(p float64) {
	l.power = p
}
