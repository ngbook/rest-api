package handler

/**
 * Author: jsongo<jsongo@qq.com>
 */

import (
	"fmt"
	"log"

	"github.com/ngbook/micro-util/net"

	proto "github.com/ngbook/rest-api/people/proto"

	"gopkg.in/mgo.v2"

	restful "github.com/emicklei/go-restful"
)

// PeopleAPI is the pos api stuct
type PeopleAPI struct {
}

// GetList is the api for list all the needs
func (wx *PeopleAPI) GetList(req *restful.Request, rsp *restful.Response) {
	log.Print("Received PeopleAPI.GetList API request")

	var status int64 = -1
	var msg interface{}

	ip := req.Request.Header.Get("X-Real-IP")
	fmt.Println(ip)

	// get param
	listReq := new(proto.GetListReq)
	// read request data
	err := req.ReadEntity(listReq)
	fmt.Println(listReq)
	if err != nil {
		msg = "invalid post data"
	} else {
		// mongo
		session, err := mgo.Dial("mongodb://mongo.techzto.svc.cluster.local")
		if err != nil {
			msg = err.Error()
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB("ngbook").C("people")
		result := []proto.People{}
		start := int(listReq.Start)
		limit := int(listReq.PageSize)
		err = c.Find(nil).Skip(start).Limit(limit).All(&result)
		if err != nil {
			msg = err.Error()
		} else {
			status = 200
			data := map[string]interface{}{
				"list":  result,
				"start": start,
				"end":   start + len(result),
				"total": 0,
			}

			// get total
			total, err := c.Count()
			if err != nil {
				msg = err.Error()
			} else {
				data["total"] = total
				msg = data
			}
		}
	}
	net.InitRsp(status, rsp, msg)
}
