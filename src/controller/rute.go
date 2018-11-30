package controller

import (
	"config"
	"fmt"
	"lib/tpe"
	"lib/unit"
	"net/http"
	"reflect"
)

//路由分发
func RuteRun(reponse http.ResponseWriter, request *http.Request) {
	//query := request.URL.Query()
	//act := query.Get("f")
	uri := request.RequestURI
	rute, _ := unit.ParseURI(uri)
	ruteExit := false
	var rcla interface{}
	switch rute[0] {
	case "Rute":
		rcla = &Rute{rp: reponse, rq: request}
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

//基类
type Rute struct {
	rp http.ResponseWriter
	rq *http.Request
}

func (obj *Rute) Index() {
	fmt.Fprint(obj.rp, "rute index")
}

func (obj *Rute) Icon() {
	obj.rp.Header().Set("Content-Type", "text/html")
	//调用模版
	view := tpe.Assign("index/icon.html")
	locals := make(map[string]interface{})
	locals["title"] = "Admin 3.0"
	locals["info"] = []string{}
	view.Execute(obj.rp, locals)
}

func (obj *Rute) Table() {
	mysql := config.DbGman()
	r := mysql.GetRow("SHOW TABLES")
	fmt.Fprint(obj.rp, r)
}
