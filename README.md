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

Output result:
```
go run cmd/main.go 
--------- a.Prop 11111
8
--------- a.Prop 22222
23

```