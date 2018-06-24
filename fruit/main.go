package main

/**
 * Author: jsongo<jsongo@qq.com>
 */

import (
	"log"

	"github.com/ngbook/micro-util/route"

	"github.com/ngbook/rest-api/fruit/handler"

	"github.com/micro/go-web"
	k8s "github.com/micro/kubernetes/go/web"
)

var (
	srvName = "net.ngbook.api.fruit"
)

func main() {

	// Create service
	service := k8s.NewService(
		web.Name(srvName),
	)

	service.Init()

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
