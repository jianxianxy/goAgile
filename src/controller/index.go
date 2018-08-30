package controller

import (
    "fmt"
    "lib/tpe"
	"net/http"
    "encoding/json"  
)


//首页
func IndexIndex(rp http.ResponseWriter, rq *http.Request) {
	rp.Header().Set("Content-Type", "text/html")
	//调用模版
	view := tpe.Assign("index/index.html");
	locals := make(map[string]interface{})
    locals["title"] = "Admin 3.0"
	locals["info"] = []string{}
	view.Execute(rp, locals)
}

//搜索 输出 json
func IndexSearch(rp http.ResponseWriter, rq *http.Request) { 
	type Road struct {  
        Name   string  
        Number int  
    }  
    roads := []Road{  
        {"Diamond Fork", 29},  
        {"Sheep Creek", 51},  
    }  
  
    ret, _ := json.Marshal(roads)
    fmt.Fprint(rp, string(ret))
}
