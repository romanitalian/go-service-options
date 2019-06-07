# go-service-options

based on article: https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html?m=1


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
	fmt.Println("## Variant1:")
	variant1()
	fmt.Println()

	fmt.Println("## Variant2:")
	variant2()
	fmt.Println()

	fmt.Println("## Variant3:")
	variant3()
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

func variant3() {
	foo := subscriptions.New()

	srvAccounts := accounts.New(8)
	srvBilling := billing.New()

	ac := subscriptions.SrvAccounts(srvAccounts)
	bl := subscriptions.SrvBilling(srvBilling)

	foo.Option(ac)
	foo.Option(bl)

	foo.Process(23)
}

```


Output result:
```
## Variant1:
--------- a.Prop 11111
8
--------- a.Prop 22222
23

## Variant2:
--------- a.Prop 11111
8
--------- a.Prop 22222
23

## Variant3:
--------- a.Prop 11111
8
--------- a.Prop 22222
23

```