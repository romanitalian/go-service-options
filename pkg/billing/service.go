package billing

// New initialize billing struct
func New() Client {
	return &billing{}
}

// Client is interface for billing struct
type Client interface {
	Increase()
}

type billing struct {
	balance string
}

func (a billing) Increase() {
	// do stuff
}
