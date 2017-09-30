package controller

import (
	"net/http"
    "fmt"
)
func Demo(reponse http.ResponseWriter, request *http.Request){
    claDemo := &ClassDemo{rp:reponse,rq:request}
    claDemo.Index();
}

type ClassDemo struct{
    rp http.ResponseWriter
    rq *http.Request    
}

func (obj *ClassDemo) Index(){
    fmt.Fprint(obj.rp,"OK")    
}