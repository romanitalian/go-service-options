package accounts

import "fmt"

func New(propVal int64) Client {
	return &accounts{Prop: propVal}
}

type Client interface {
	Bind(int64)
}

type accounts struct {
	Prop int64
}

func (a accounts) Bind(userId int64) {
	fmt.Println("--------- a.Prop 11111")
	fmt.Println(a.Prop)

	a.Prop = userId
	fmt.Println("--------- a.Prop 22222")
	fmt.Println(a.Prop)
}