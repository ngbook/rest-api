package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	proto "github.com/ngbook/rest-api/fruit/proto"

	"github.com/ngbook/micro-util/net"

	"gopkg.in/mgo.v2"

	restful "github.com/emicklei/go-restful"
)

// FruitAPI is the pos api stuct
type FruitAPI struct {
}

// GetList is the api for list all the needs
func (wx *FruitAPI) GetList(req *restful.Request, rsp *restful.Response) {
	log.Print("Received FruitAPI.GetList API request")

	var status int64 = -1
	var msg interface{}

	ip := req.Request.Header.Get("X-Real-IP")
	fmt.Println(ip)
	// mongo
	session, err := mgo.Dial("mongodb://mongo.techzto.svc.cluster.local")
	if err != nil {
		msg = err.Error()
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("ngbook").C("fruit")
	result := []proto.Fruit{}
	err = c.Find(nil).All(&result)
	if err != nil {
		msg = err.Error()
	} else {
		status = 200
		msg = result
	}
	net.InitRsp(status, rsp, msg)
}

// RspError is the api for list all the needs
func (wx *FruitAPI) RspError(req *restful.Request, rsp *restful.Response) {
	log.Print("Received FruitAPI.RspError API request")

	codeStr := req.PathParameter("code")
	var (
		code int64 = 1001
		msg  string
	)
	fmt.Println(codeStr)
	if len(codeStr) != 0 {
		tmpCode, err := strconv.Atoi(codeStr)
		fmt.Println(tmpCode)
		if err != nil {
			msg = "/error/路径后面只能传入数字"
		} else {
			code = int64(tmpCode)
		}
	}
	if len(msg) == 0 {
		msg = fmt.Sprintf("模拟%d错误", code)
	}

	result, err := json.Marshal(map[string]interface{}{
		"code": 1001,
		"msg":  "模拟500错误",
	})
	if err != nil {
		rsp.WriteErrorString(500, "{}")
	}
	rsp.AddHeader("Content-Type", "application/json; charset=utf-8")
	rsp.WriteErrorString(500, string(result))
	// rsp.WriteError(500, errors.New("模拟500错误"))
}

// RspErrorCode is the api for list all the needs
func (wx *FruitAPI) RspErrorCode(req *restful.Request, rsp *restful.Response) {
	log.Print("Received FruitAPI.RspErrorCode API request")
	codeStr := req.PathParameter("code")
	var (
		code int64 = 1001
		msg  string
	)
	fmt.Println(codeStr)
	if len(codeStr) != 0 {
		tmpCode, err := strconv.Atoi(codeStr)
		fmt.Println(tmpCode)
		if err != nil {
			msg = "/error/路径后面只能传入数字"
		} else {
			code = int64(tmpCode)
		}
	}
	if len(msg) == 0 {
		msg = fmt.Sprintf("模拟%d错误", code)
	}

	net.InitRsp(code, rsp, msg)
}
