package controller

import (
	"lib/tpe"
	"net/http"
)

//基类
type Index struct {
	rp http.ResponseWriter
	rq *http.Request
}

//首页
func (obj *Index) Index() {
	obj.rp.Header().Set("Content-Type", "text/html")
	//调用模版
	view := tpe.Assign("index/index.html")
	locals := make(map[string]interface{})
	locals["title"] = "Admin 3.0"
	locals["info"] = []string{}
	view.Execute(obj.rp, locals)
}

//图标
func (obj *Index) Icon() {
	obj.rp.Header().Set("Content-Type", "text/html")
	//调用模版
	view := tpe.Assign("index/icon.html")
	locals := make(map[string]interface{})
	locals["title"] = "Admin 3.0"
	locals["info"] = []string{}
	view.Execute(obj.rp, locals)
}
