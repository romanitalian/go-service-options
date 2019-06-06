package billing

func New() Client {
	return &billing{}
}

type Client interface {
	Increase()
}

type billing struct {
	balance string
}

func (a billing) Increase() {
	// do stuff
}
