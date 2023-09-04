package component

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var _ Componenter = (*EnvComponent)(nil)

func NewEnvComponent(name string, cType, port int) Componenter {
	return &EnvComponent{
		Name:    name,
		payload: []*EnvPayload{},
		Type:    []int{cType},
		port:    port,
	}
}

func (e *EnvComponent) Register(h *Huber) {
	e.hubers = append(e.hubers, h)
}

func (e *EnvComponent) Recieve(data ReceivePayloader) {
	// 型アサーション
	sData, ok := data.(EnvPayload)
	if !ok {
		log.Printf("Failed to type assert data: %v to EnvPayload", data)
		return
	}
	log.Printf("Recieved Temperature: %f", sData.Temperature)
	e.SetPayload(NewEnvPayload(sData.Temperature))
}

func (e *EnvComponent) GetType() []int {
	return e.Type
}

func (e *EnvComponent) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/env", e.GetFromHttp)
	log.Printf("Start %v Server. port:%d", e.Name, e.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", e.port), mux)
}

func (e *EnvComponent) GetFromHttp(w http.ResponseWriter, r *http.Request) {
	temperature, err := strconv.ParseFloat(r.URL.Query().Get("t"), 64)
	if err != nil {
		http.Error(w, "Invalid env value", http.StatusBadRequest)
		return
	}
	e.SetPayload(NewEnvPayload(temperature))
	e.Notify()
}

func (e *EnvComponent) Notify() {
	p := e.payload[0]
	e.payload = e.payload[1:]

	payload := NewReceivePayload(EnvEventFromBack, p)
	for _, ob := range e.hubers {
		go (*ob).Receive(payload)
	}
}

func (e *EnvComponent) SetPayload(ep *EnvPayload) {
	e.payload = append(e.payload, ep)
}
