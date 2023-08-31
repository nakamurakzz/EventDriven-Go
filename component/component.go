package component

type Componenter interface {
	Register(h *Huber)
	Notify()
	Recieve(data interface{})
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

type EnvPayload struct {
	Temperature float64
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

type ReceivePayload struct {
	eventType int
	data      interface{}
}

type Hub struct {
	observers map[int][]Componenter
}

type LightComponent struct {
	Name    string
	payload []*LightPayload
	Type    []int
	hubers  []*Huber
	port    int
}
