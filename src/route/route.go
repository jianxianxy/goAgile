package route

import (
	"controller"
	"fmt"
	"lib/unit"
	"net/http"
	"reflect"
)

//注册路由
func GetRouteMap(reponse http.ResponseWriter, request *http.Request) (RouteMap map[string]interface{}) {
	RouteMap = make(map[string]interface{})
	RouteMap["Index"] = &controller.Index{Rp: reponse, Rq: request}
	RouteMap["Data"] = &controller.Data{Rp: reponse, Rq: request}
	return
}

//路由分发
func RouteRun(reponse http.ResponseWriter, request *http.Request) {
	uri := request.RequestURI
	route, _ := unit.ParseURI(uri)
	clact := route[0] //默认Index
	/*
			ruteExit := false
			var rcla interface{}
			switch route[0] {
			case "Index":
				rcla = &Index{rp: reponse, rq: request}
			case "Data":
				rcla = &Data{rp: reponse, rq: request}
			default:
				ruteExit = true
				fmt.Println("Route Error!", route)
			}
		if ruteExit == false {
			rftObj := reflect.ValueOf(rcla).MethodByName(route[1])
			if rftObj.IsValid() == false {
				fmt.Println("Route model Error!", route)
			} else {
				param := make([]reflect.Value, rftObj.Type().NumIn())
				rftObj.Call(param)
			}
		}
	*/
	RouteMap := GetRouteMap(reponse, request)
	if rcla, ok := RouteMap[clact]; ok {
		rftObj := reflect.ValueOf(rcla).MethodByName(route[1])
		if rftObj.IsValid() == false {
			fmt.Println("Route model Error!", route)
		} else {
			param := make([]reflect.Value, rftObj.Type().NumIn())
			rftObj.Call(param)
		}
	} else {
		fmt.Println("Route model null!", route)
	}
}
