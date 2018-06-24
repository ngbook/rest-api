package main

/**
 * Author: jsongo<jsongo@qq.com>
 */

import (
	"log"

	"micro/common/route"

	"micro/ngbook/people/rest/handler"

	"github.com/micro/go-micro/registry"
	web "github.com/micro/go-web"
)

var (
	srvName = "net.ngbook.api.people"
	addr    = "consul.techzto.com:80"
)

func main() {

	reg := registry.NewRegistry(registry.Addrs(addr))
	// Create service
	service := web.NewService(
		web.Name(srvName),
		web.Registry(reg),
	)

	service.Init()

	// Create RESTful handler
	peopleAPI := new(handler.PeopleAPI)
	// router
	router := route.NewRouter().AddRouter(
		"/people", route.RouterOption{
			" <= post": route.HandleOption{
				Handler: peopleAPI.GetList,
			},
		},
	)
	service.Handle("/", router.WS)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
