package unit

import(
    "io/ioutil"
    "regexp"
    "strings"
    "strconv"
)

//返回视图参数
func GetParam(file,mod string) []map[string]string{
    con := ReadFile(file)
    match := FindCol(con)
    res := ConfigMap(match,mod)
    return res
}

//读取文件内容
func ReadFile(file string) string{
    con,err := ioutil.ReadFile(file)
    if err == nil{
        return string(con)
    }else{
        return ""
    }
}

//正则解析内容
func FindCol(str string) []string{
    //清除空白
    repExg,_ := regexp.Compile(`\s`)
    str = repExg.ReplaceAllString(str, "")
    //匹配正文
    reg := regexp.MustCompile(`\[[^]]*\]`)
    match := reg.FindAllString(str, -1)
    return match
}

//生成Map
func ConfigMap(data []string,mod string) []map[string]string{
    count := len(data)
    colMap := make([]map[string]string,count)
    for i := range data{
        repExg,_ := regexp.Compile(`\[|\]`)
        str := repExg.ReplaceAllString(data[i], "")
        kv := strings.Split(str,",")
        tmp := make(map[string]string)
        for j := range kv{
            kvarr := strings.Split(kv[j],":")
            tmp[kvarr[0]] = kvarr[1]
        }
        index,_ := strconv.Atoi(tmp[mod])
        colMap[index] = tmp
    }
    var end int
    for i := range colMap{
        if len(colMap[i]) == 0{
            end = i
            break    
        }    
    }
    if end > 1{
        return colMap[1:end]    
    }else{
        return colMap[1:]
    }
}