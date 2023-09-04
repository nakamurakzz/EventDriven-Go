package component

type EventType int

// EventTypes
const (
	LightEventFromBack = iota
	EnvEventFromBack
	LightEventFromFront
	EnvEventFromFront
	LightEvent
	EnvEvent
)

type Componenter interface {
	Register(h *Huber)
	Notify()
	Recieve(data ReceivePayloader)
	GetType() []int
	Start() error
}

func InitializeComponent() []Componenter {
	return []Componenter{
		NewEnvComponent("EnvFront", EnvEventFromBack, 6001),
		NewEnvComponent("EnvBack", EnvEventFromFront, 6002),
		NewLightComponent("LightFront", LightEventFromBack, 7001),
		NewLightComponent("LightBack", LightEventFromFront, 7002),
	}
}

type Huber interface {
	Notify(payload ReceivePayload)
	Receive(payload ReceivePayload)
	Register(o Componenter)
}

type LightPayload struct {
	Power float64
}

func NewLightPayload(p float64) *LightPayload {
	return &LightPayload{
		Power: p,
	}
}

func (l LightPayload) GetType() int {
	return LightEventFromBack
}

type EnvPayload struct {
	Temperature float64
}

func (e EnvPayload) GetType() int {
	return EnvEventFromBack
}

func NewEnvPayload(t float64) *EnvPayload {
	return &EnvPayload{
		Temperature: t,
	}
}

type EnvComponent struct {
	Name    string
	payload []*EnvPayload
	Type    []int
	hubers  []*Huber
	port    int // 暫定, Configで管理する
}

type ReceivePayloader interface {
	GetType() int
}

func ReceivePayloads[T ReceivePayloader]() []T {
	return nil
}

type ReceivePayload struct {
	eventType EventType
	data      ReceivePayloader
}

func (r ReceivePayload) GetType() EventType {
	return r.eventType
}

type Hub struct {
	observers map[EventType][]Componenter
}

type LightComponent struct {
	Name    string
	payload []*LightPayload
	Type    []int
	hubers  []*Huber
	port    int
}
