package controller

import (
    "fmt"
    "config"
	"net/http"
    "encoding/json"  
)

//路由分发
func DemoRun(reponse http.ResponseWriter, request *http.Request){
    claDemo := &Demo{rp:reponse,rq:request}
    query := request.URL.Query()
    act := query.Get("f")    
    switch act{
        case "index":
            claDemo.index()
        case "json":
            claDemo.json()
        case "table":
            claDemo.table()
        default:
            claDemo.index()    
    }
    
}
//基类
type Demo struct{
    rp http.ResponseWriter
    rq *http.Request
}

func (obj *Demo) index(){
    fmt.Fprint(obj.rp,"index")
}

func (obj *Demo) json() {
	type Road struct {  
        Name   string  
        Number int  
    }  
    roads := []Road{  
        {"Diamond Fork", 29},  
        {"Sheep Creek", 51},  
    }  
    ret, _ := json.Marshal(roads)
    fmt.Fprint(obj.rp, string(ret))
}

func (obj *Demo) table(){
    mysql := config.DbGman()
    r := mysql.GetRow("SHOW TABLES")
    fmt.Fprint(obj.rp, r)
}