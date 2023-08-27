package sendor

type Sendor interface {
	Recieve(data interface{})
	Print()
	GetSendorType() string
}
