package component

type Componenter interface {
	Register(h *Huber)
	Notify()
	Recieve(data interface{})
	GetType() []int
	Start() error
}

type LightPayload struct {
	Power float64
}

type EnvPayload struct {
	Temperature float64
}

func InitializeComponent() []Componenter {
	return []Componenter{
		NewEnvBackComponent(),
		NewLightBackComponent(),
		NewEnvFrontComponent(),
		NewLightFrontComponent(),
	}
}

type EnvBackComponent struct {
	temperature float64
	Type        []int
	hubers      []*Huber
}

type EnvFrontComponent struct {
	temperature float64
	Type        []int
	hubers      []*Huber
}

type ReceivePayload struct {
	eventType int
	data      interface{}
}

type Huber interface {
	Notify(payload ReceivePayload)
	Receive(payload ReceivePayload)
	Register(o Componenter)
}

type Hub struct {
	observers map[int][]Componenter
}

type LightBackComponent struct {
	power  float64
	Type   []int
	hubers []*Huber
}

type LightFrontComponent struct {
	power  float64
	Type   []int
	hubers []*Huber
}
