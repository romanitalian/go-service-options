package main

import (
	"fmt"

	"go-service-options/pkg/accounts"
	"go-service-options/pkg/billing"
	"go-service-options/pkg/subscriptions"
)

func main() {
	fmt.Println("Variant_1:")
	variant1()

	fmt.Println("Variant_2:")
	variant2()

	fmt.Println("Variant_3:")
	variant3()
}


func variant1() {
	accs := accounts.New()

	srvSubscriptions := subscriptions.New(111)

	sbs := accounts.BindSubscriptions(srvSubscriptions)

	accs.AddOptions(sbs)

	accs.Process(222)
}

func variant2() {
	accs := accounts.New()

	srvSubscriptions := subscriptions.New(333)
	srvBilling := billing.New()

	sbs := accounts.BindSubscriptions(srvSubscriptions)
	bl := accounts.BindBilling(srvBilling)

	accs.AddOptions(sbs, bl)

	accs.Process(444)
}

func variant3() {
	accs := accounts.New()

	srvSubscriptions := subscriptions.New(555)
	srvBilling := billing.New()

	sbs := accounts.BindSubscriptions(srvSubscriptions)
	bl := accounts.BindBilling(srvBilling)

	accs.AddOptions(sbs)
	accs.AddOptions(bl)

	accs.Process(777)
}
