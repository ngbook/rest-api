package main

/**
 * Author: jsongo<jsongo@qq.com>
 */

import (
	"log"

	"github.com/ngbook/micro-util/route"

	"github.com/ngbook/rest-api/people/handler"

	web "github.com/micro/go-web"
)

var (
	srvName = "net.ngbook.api.people"
)

func main() {

	// Create service
	service := web.NewService(
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
