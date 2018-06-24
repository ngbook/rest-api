package main

/**
 * Author: jsongo<jsongo@qq.com>
 */

import (
	"log"

	"github.com/ngbook/micro-util/route"

	"github.com/ngbook/rest-api/people/handler"

	"github.com/micro/go-web"
	k8s "github.com/micro/kubernetes/go/web"
)

var (
	srvName = "net.ngbook.api.people"
)

func main() {

	// Create service
	service := k8s.NewService(
		web.Name(srvName),
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
