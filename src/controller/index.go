package controller

import (
	"config"
	"html/template"
	"net/http"
    "encoding/json"  
    "fmt"
)


//首页
func Index(rp http.ResponseWriter, rq *http.Request) {
	rp.Header().Set("Content-Type", "text/html")
	//调用模版
	view, err := template.ParseFiles(config.Get("ROOT_PATH") + "static/view/index.html")
	if err != nil {
		http.Error(rp, err.Error(), http.StatusInternalServerError)
		return
	}
	locals := make(map[string]interface{})
	locals["info"] = []string{}
	view.Execute(rp, locals)
}

//搜索 输出 json
func Search(rp http.ResponseWriter, rq *http.Request) {

    fmt.Println(config.GetGmanConf())
    
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
