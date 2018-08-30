package main

import(
    "fmt"
    "controller"
    "net/http"
)

func main(){
    
    //启动静态文件服务
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    
    routeHandler()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("启动失败: ", err.Error())
	}
}

//路由
func routeHandler() {
	http.HandleFunc("/", controller.IndexRun)
	http.HandleFunc("/Index", controller.IndexRun)
    http.HandleFunc("/Demo", controller.DemoRun)
}