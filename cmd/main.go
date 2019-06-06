package main

import (
	"../pkg/accounts"
	"../pkg/billing"
	"../pkg/subscriptions"
)

func main() {
	variant1()
	variant2()
}


func variant1() {
	foo := subscriptions.Foo{}

	srvAccounts := accounts.New(8)

	ac := subscriptions.SrvAccounts(srvAccounts)

	foo.Option(ac)

	foo.Process(23)
}

func variant2() {
	foo := subscriptions.Foo{}

	srvAccounts := accounts.New(8)
	srvBilling := billing.New()

	ac := subscriptions.SrvAccounts(srvAccounts)
	bl := subscriptions.SrvBilling(srvBilling)

	foo.Option(ac, bl)

	foo.Process(23)
}
