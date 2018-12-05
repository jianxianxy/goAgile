package unit

import (
	"io/ioutil"
	"strings"
)

//读取文件内容
func ReadFile(file string) string {
	con, err := ioutil.ReadFile(file)
	if err == nil {
		return string(con)
	} else {
		return ""
	}
}

//解析URL获取路由信息
func ParseURI(url string) (rute [2]string, num int) {
	rute = [2]string{"Index", "Index"}
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
