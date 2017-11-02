package controller

import (
	"net/http"
    "fmt"
)

//基类
type ClassDemo struct{
    rp http.ResponseWriter
    rq *http.Request    
}

//路由分发
func Demo(reponse http.ResponseWriter, request *http.Request){
    claDemo := &ClassDemo{rp:reponse,rq:request}
    query := request.URL.Query()
    act := query.Get("f")
    switch act{
        case "index":
            claDemo.index()
        case "demo":
            claDemo.demo()
        default:
            claDemo.index()    
    }
    
}

func (obj *ClassDemo) index(){
    fmt.Fprint(obj.rp, "index")
}

func (obj *ClassDemo) demo() {
    fmt.Fprint(obj.rp, "demo")
}