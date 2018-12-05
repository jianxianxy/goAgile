package controller

import (
	"config"
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

//统计图
func (obj *Data) ChartJson() {
	month := ReqGet("month", obj.rq)
	data := make(map[string]interface{})
	sel := "SELECT *,DATE_FORMAT(`anday`,'%m.%d') AS `mday` FROM `data_chart` WHERE DATE_FORMAT(`anday`,'%Y-%m') = '" + month + "' ORDER BY anday ASC"
	mysql := config.DbSpider()
	all := mysql.GetRow(sel)
	xAxis := make([]string, 0)
	yAdd := make([]string, 0)
	yReduce := make([]string, 0)
	if len(all) > 0 {
		for _, val := range all {
			xAxis = append(xAxis, val["mday"])
			yAdd = append(yAdd, val["add"])
			yReduce = append(yReduce, val["reduce"])
		}
		data["xAxis"] = xAxis
		//平均线
		markLine := make(map[string]interface{})
		var lineCol = map[string]string{"type": "average", "name": "平均值"}
		markLine["data"] = []map[string]string{lineCol}
		//数据
		var yAxis [2]interface{}
		yAxis0 := make(map[string]interface{})
		yAxis0["name"] = "新增"
		yAxis0["type"] = "bar"
		yAxis0["data"] = yAdd
		yAxis0["markLine"] = markLine
		yAxis1 := make(map[string]interface{})
		yAxis1["name"] = "销售"
		yAxis1["type"] = "bar"
		yAxis1["data"] = yReduce
		yAxis1["markLine"] = markLine

		yAxis[0] = yAxis0
		yAxis[1] = yAxis1

		data["series"] = yAxis
	}
	ret, _ := json.Marshal(data)
	fmt.Fprint(obj.rp, string(ret))
}
