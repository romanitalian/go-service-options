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
```


Output result:
```
## variant1:
--------- a.Prop 11111
8
--------- a.Prop 22222
23

## variant2:
--------- a.Prop 11111
8
--------- a.Prop 22222
23

```