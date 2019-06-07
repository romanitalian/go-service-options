package main

import (
	"../pkg/accounts"
	"../pkg/billing"
	"../pkg/subscriptions"
	"fmt"
)

func main() {
	fmt.Println("## variant1:")
	variant1()
	fmt.Println()

	fmt.Println("## variant2:")
	variant2()
}


func variant1() {
	foo := subscriptions.New()

	srvAccounts := accounts.New(8)

	ac := subscriptions.SrvAccounts(srvAccounts)

	foo.Option(ac)

	foo.Process(23)
}

func variant2() {
	foo := subscriptions.New()

	srvAccounts := accounts.New(8)
	srvBilling := billing.New()

	ac := subscriptions.SrvAccounts(srvAccounts)
	bl := subscriptions.SrvBilling(srvBilling)

	foo.Option(ac, bl)

	foo.Process(23)
}
