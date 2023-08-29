package sendor

type Sendorer interface {
	// use Generics
	Receive(data interface{})
	Print()
	GetSendorType() string
}
