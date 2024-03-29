module github.com/chremoas/esi-srv

go 1.14

require (
	github.com/antihax/goesi v0.0.0-20190723215635-487b927dc566
	github.com/chremoas/services-common v1.3.2
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79
	github.com/mailru/easyjson v0.0.0-20190626092158-b2ccc519800e // indirect
	github.com/micro/go-micro v1.9.1
	github.com/petergtz/pegomock v2.8.0+incompatible
	github.com/smartystreets/goconvey v0.0.0-20190710185942-9d28bd7c0945
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
)

replace github.com/chremoas/esi-srv => ../esi-srv

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
