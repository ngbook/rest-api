package main

import (
	"log"

	"github.com/ngbook/micro-util/route"

	"github.com/ngbook/rest-api/fruit/handler"

	"github.com/micro/go-web"
	k8s "github.com/micro/kubernetes/go/web"
)

var (
	srvName = "net.ngbook.api.fruit"
	// targetServer = "com.jsongo.ngbook.srv.pos"
	// addr = "consul.techzto.com:80"
)

func main() {

	// reg := registry.NewRegistry(registry.Addrs(addr))
	// Create service
	service := k8s.NewService(
		web.Name(srvName),
	)
	// srv := &registry.Service{ Name: "go.micro.srv.auth" }
	// service := micro.NewService(
	// 	micro.Name("go.micro.srv.auth"),
	// 	micro.Version("latest"),
	// 	micro.WrapClient(trace.ClientWrapper(t, srv)),
	// 	micro.WrapHandler(trace.HandlerWrapper(t, srv)),
	// )

	service.Init()

	// setup Server Client
	// handler.InitWxClient(
	// proto.NewWechatClient(targetServer, client.DefaultClient))

	// Create RESTful handler
	fruitAPI := new(handler.FruitAPI)
	// router
	router := route.NewRouter().AddRouter(
		"/fruit", route.RouterOption{
			" <= post": route.HandleOption{
				Handler: fruitAPI.GetList,
			},
			"/error <= post": route.HandleOption{
				Handler: fruitAPI.RspError,
			},
			"/error-code/{code} <= post": route.HandleOption{
				Handler: fruitAPI.RspErrorCode,
			},
		},
	)
	service.Handle("/", router.WS)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
