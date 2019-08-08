package controller

import "net/http"

//获取Request
func ReqGet(name string, req *http.Request) string {
	query := req.URL.Query()
	act := query.Get(name)
	return act
}
