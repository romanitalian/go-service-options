package accounts

import (
	"../billing"
	"../subscriptions"
)

// Init ------------------

// New initialize account struct
func New() Client {
	return &accs{}
}

// Client is interface for account struct
type Client interface {
	AddOptions(...option)
	Process(int64)
}

type accs struct {
	verbosity        int
	srvSubscriptions subscriptions.Client
	srvBilling       billing.Client
}

// Dependencies ------------------

// Option self-referential function
type option interface {
	apply(*accs)
}

func (f *accs) AddOptions(opts ...option) {
	for _, opt := range opts {
		opt.apply(f)
	}
}

// Verbosity set verbosity value
func Verbosity(v int) Option {
	return func(f *accs) {
		f.verbosity = v
	}
}

// BindSubscriptions - bind property: subscription service
func BindSubscriptions(v subscriptions.Client) Option {
	return func(f *accs) {
		f.srvSubscriptions = v
	}
}

// BindBilling - bind property: billing service
func BindBilling(v billing.Client) Option {
	return func(f *accs) {
		f.srvBilling = v
	}
}

// Public methods ------------------

func (f accs) Process(val int64) {
	f.srvSubscriptions.Bind(val)
}
