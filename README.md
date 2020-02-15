# go-service-options


```
.
├── cmd
│   └── main.go
├── pkg
│   ├── accounts
│   │   └── service.go
│   ├── billing
│   │   └── service.go
│   └── subscriptions
│       └── service.go
└── README.md
```

Entry point:
```go
package main

import (
	"../pkg/accounts"
	"../pkg/billing"
	"../pkg/subscriptions"
	"fmt"
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
	subs := subscriptions.New()

	srvAccounts := accounts.New(111)

	ac := subscriptions.BindAccounts(srvAccounts)

	subs.AddOptions(ac)

	subs.Process(222)
}

func variant2() {
	subs := subscriptions.New()

	srvAccounts := accounts.New(333)
	srvBilling := billing.New()

	ac := subscriptions.BindAccounts(srvAccounts)
	bl := subscriptions.BindBilling(srvBilling)

	subs.AddOptions(ac, bl)

	subs.Process(444)
}

func variant3() {
	subs := subscriptions.New()

	srvAccounts := accounts.New(555)
	srvBilling := billing.New()

	ac := subscriptions.BindAccounts(srvAccounts)
	bl := subscriptions.BindBilling(srvBilling)

	subs.AddOptions(ac)
	subs.AddOptions(bl)

	subs.Process(777)
}


```


Output result:
```
Variant_1:
        a.Prop (before): 111
        a.Prop (after): 222

Variant_2:
        a.Prop (before): 333
        a.Prop (after): 444

Variant_3:
        a.Prop (before): 555
        a.Prop (after): 777

```