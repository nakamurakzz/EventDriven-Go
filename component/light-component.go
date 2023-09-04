package component

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var _ Componenter = (*LightComponent)(nil)

func NewLightComponent(name string, cType, port int) *LightComponent {
	return &LightComponent{
		Name:    name,
		payload: []*LightPayload{},
		Type:    []int{cType},
		port:    port,
	}
}

func (l *LightComponent) Register(h *Huber) {
	l.hubers = append(l.hubers, h)
}

func (l *LightComponent) Recieve(data ReceivePayloader) {
	evType := data.GetType()
	if evType != LightEventFromFront {
		log.Printf("Failed to type assert data: %v to LightPayload", data)
		return
	}
	log.Printf("Recieved Power: %f", data.Power)
	l.SetPayload(NewLightPayload(sData.Power))
}

func (l *LightComponent) GetType() []int {
	return l.Type
}

func (l *LightComponent) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/light", l.getFromHttp)
	log.Printf("Start %v Server. port:%d", l.Name, l.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", l.port), mux)
}

func (l *LightComponent) getFromHttp(w http.ResponseWriter, r *http.Request) {
	power, err := strconv.ParseFloat(r.URL.Query().Get("p"), 64)
	if err != nil {
		http.Error(w, "Invalid power value", http.StatusBadRequest)
		return
	}
	l.SetPayload(NewLightPayload(power))
	l.Notify()
}

func (l *LightComponent) Notify() {
	p := l.payload[0]
	l.payload = l.payload[1:]

	payload := NewReceivePayload(LightEventFromBack, p)

	for _, ob := range l.hubers {
		go (*ob).Receive(payload)
	}
}

func (e *LightComponent) SetPayload(lp *LightPayload) {
	e.payload = append(e.payload, lp)
}
