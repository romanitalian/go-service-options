package subscriptions

import "fmt"

// New initialize subscriptions struct
func New(propVal int64) Client {
	return &subs{Prop: propVal}
}

// Client is interface for subscriptions struct
type Client interface {
	Bind(int64)
}

type subs struct {
	Prop int64
}

func (s subs) Bind(propValue int64) {
	fmt.Println(fmt.Sprintf("\ts.Prop (before): %#v", s.Prop))

	s.Prop = propValue
	fmt.Println(fmt.Sprintf("\ts.Prop (after): %#v\n", s.Prop))
}