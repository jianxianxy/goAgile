package controller

import (
	"encoding/json"
	"fmt"
	"lib/tpe"
	"net/http"
)

//基类
type Data struct {
	rp http.ResponseWriter
	rq *http.Request
}

//首页
func (obj *Data) Index() {
	obj.rp.Header().Set("Content-Type", "text/html")
	//调用模版
	view := tpe.Assign("data/index.html")
	locals := make(map[string]interface{})
	locals["title"] = "数据分析"
	locals["info"] = []string{}
	view.Execute(obj.rp, locals)
}

//搜索 输出 json
func (obj *Data) Json() {
	data := make(map[string]interface{})
	xAxis := [12]string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"}
	data["xAxis"] = xAxis
	y1 := [12]float32{2.0, 4.9, 7.0, 23.2, 25.6, 76.7, 135.6, 162.2, 32.6, 20.0, 6.4, 3.3}
	y2 := [12]float32{2.6, 5.9, 9.0, 26.4, 28.7, 70.7, 175.6, 182.2, 48.7, 18.8, 6.0, 2.3}

	markLine := make(map[string]interface{})
	var lineCol = map[string]string{"type": "average", "name": "平均值"}
	markLine["data"] = []map[string]string{lineCol}
	var yAxis [2]interface{}
	yAxis0 := make(map[string]interface{})
	yAxis0["name"] = "蒸发量"
	yAxis0["type"] = "bar"
	yAxis0["data"] = y1
	yAxis0["markLine"] = markLine
	yAxis1 := make(map[string]interface{})
	yAxis1["name"] = "降水量"
	yAxis1["type"] = "bar"
	yAxis1["data"] = y2
	yAxis1["markLine"] = markLine

	yAxis[0] = yAxis0
	yAxis[1] = yAxis1

	data["series"] = yAxis
	ret, _ := json.Marshal(data)
	fmt.Fprint(obj.rp, string(ret))
}
