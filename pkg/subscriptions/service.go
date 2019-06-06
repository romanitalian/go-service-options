package subscriptions

import (
	"../accounts"
	"../billing"
)

type Foo struct {
	verbosity int
	srvAccounts accounts.Client
	srvBilling billing.Client
}
type Option func(*Foo)

func (f *Foo) Option(opts ...Option) {
	for _, opt := range opts {
		opt(f)
	}
}

func Verbosity(v int) Option {
	return func(f *Foo) {
		f.verbosity = v
	}
}

func SrvAccounts(v accounts.Client) Option {
	return func(f *Foo) {
		f.srvAccounts = v
	}
}

func SrvBilling(v billing.Client) Option {
	return func(f *Foo) {
		f.srvBilling = v
	}
}

func (f Foo) Process(val int64) {
	f.srvAccounts.Bind(val)
}
