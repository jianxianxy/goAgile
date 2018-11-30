package unit

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

//返回视图参数
func GetParam(file, mod string) []map[string]string {
	con := ReadFile(file)
	match := FindCol(con)
	res := ConfigMap(match, mod)
	return res
}

//读取文件内容
func ReadFile(file string) string {
	con, err := ioutil.ReadFile(file)
	if err == nil {
		return string(con)
	} else {
		return ""
	}
}

//正则解析内容
func FindCol(str string) []string {
	//清除空白
	repExg, _ := regexp.Compile(`\s`)
	str = repExg.ReplaceAllString(str, "")
	//匹配正文
	reg := regexp.MustCompile(`\[[^]]*\]`)
	match := reg.FindAllString(str, -1)
	return match
}

//生成Map
func ConfigMap(data []string, mod string) []map[string]string {
	count := len(data)
	colMap := make([]map[string]string, count)
	for i := range data {
		repExg, _ := regexp.Compile(`\[|\]`)
		str := repExg.ReplaceAllString(data[i], "")
		kv := strings.Split(str, ",")
		tmp := make(map[string]string)
		for j := range kv {
			kvarr := strings.Split(kv[j], ":")
			tmp[kvarr[0]] = kvarr[1]
		}
		index, _ := strconv.Atoi(tmp[mod])
		colMap[index] = tmp
	}
	var end int
	for i := range colMap {
		if len(colMap[i]) == 0 {
			end = i
			break
		}
	}
	if end > 1 {
		return colMap[1:end]
	} else {
		return colMap[1:]
	}
}

//获取表单额外参数
func GetExtData(col []map[string]string) map[string]map[string]string {
	retForm := make(map[string]map[string]string)
	for _, v := range col {
		data, have := v["data"]
		if have {
			retForm[v["col"]] = FmtRadio(data)
		}
	}
	return retForm
}

//Radio
func FmtRadio(data string) map[string]string {
	repExg, _ := regexp.Compile(`\(|\)`)
	str := repExg.ReplaceAllString(data, "")
	kv := strings.Split(str, "|")
	ret := make(map[string]string)
	for j := range kv {
		kvarr := strings.Split(kv[j], "=")
		ret[kvarr[0]] = kvarr[1]
	}
	return ret
}

//解析URL获取路由信息
func ParseURI(url string) (rute [2]string, num int) {
	rute = [2]string{"Rute", "Index"}
	urif1 := strings.Split(strings.Trim(url, "/"), "&")
	urif2 := strings.Split(urif1[0], "?")
	uriArr := strings.Split(urif2[0], "/")
	if len(uriArr[0]) < 1 {
		return rute, 0
	} else if len(uriArr) == 1 {
		rute[0] = Capitalize(uriArr[0])
	} else {
		rute[0] = Capitalize(uriArr[0])
		rute[1] = Capitalize(uriArr[1])
	}
	return rute, 2
}

//首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
