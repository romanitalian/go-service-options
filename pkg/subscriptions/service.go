package subscriptions

import (
	"../accounts"
	"../billing"
)

// Init ------------------

func New() Client {
	return &subscriptions{}
}

type Client interface {
	Option(...Option)
	Process(int64)
}

type subscriptions struct {
	verbosity int
	srvAccounts accounts.Client
	srvBilling billing.Client
}

// Dependencies ------------------

type Option func(*subscriptions)

func (f *subscriptions) Option(opts ...Option) {
	for _, opt := range opts {
		opt(f)
	}
}

func Verbosity(v int) Option {
	return func(f *subscriptions) {
		f.verbosity = v
	}
}

func SrvAccounts(v accounts.Client) Option {
	return func(f *subscriptions) {
		f.srvAccounts = v
	}
}

func SrvBilling(v billing.Client) Option {
	return func(f *subscriptions) {
		f.srvBilling = v
	}
}

// Public methods ------------------

func (f subscriptions) Process(val int64) {
	f.srvAccounts.Bind(val)
}
