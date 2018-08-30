package controller

import (
    "fmt"
    "lib/tpe"
	"net/http"
    "encoding/json"  
)

//路由分发
func IndexRun(reponse http.ResponseWriter, request *http.Request){
    claIndex := &Index{rp:reponse,rq:request}
    query := request.URL.Query()
    act := query.Get("f")    
    switch act{
        case "index":
            claIndex.index()
        case "json":
            claIndex.json()
        default:
            claIndex.index()    
    }
    
}
//基类
type Index struct{
    rp http.ResponseWriter
    rq *http.Request
}

//首页
func (obj *Index) index(){
	obj.rp.Header().Set("Content-Type", "text/html")
	//调用模版
	view := tpe.Assign("index/index.html");
	locals := make(map[string]interface{})
    locals["title"] = "Admin 3.0"
	locals["info"] = []string{}
	view.Execute(obj.rp, locals)
}

//搜索 输出 json
func (obj *Index) json() { 
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
