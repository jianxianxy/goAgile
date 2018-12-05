package controller

import (
	"fmt"
	"lib/unit"
	"net/http"
	"reflect"
)

//路由分发
func RuteRun(reponse http.ResponseWriter, request *http.Request) {
	uri := request.RequestURI
	rute, _ := unit.ParseURI(uri)
	ruteExit := false
	var rcla interface{}
	switch rute[0] {
	case "Index":
		rcla = &Index{rp: reponse, rq: request}
	case "Data":
		rcla = &Data{rp: reponse, rq: request}
	default:
		ruteExit = true
		fmt.Println("Route Error!", rute)
	}
	if ruteExit == false {
		rftObj := reflect.ValueOf(rcla).MethodByName(rute[1])
		if rftObj.IsValid() == false {
			fmt.Println("Route model Error!", rute)
		} else {
			param := make([]reflect.Value, rftObj.Type().NumIn())
			rftObj.Call(param)
		}
	}
}

//获取Request
func ReqGet(name string, req *http.Request) string {
	query := req.URL.Query()
	act := query.Get(name)
	return act
}
