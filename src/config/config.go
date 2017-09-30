/**
    配置信息
    定义的配置信息使用小写字母，保持对外私有
*/
package config

var Config = map[string]string{
	"ROOT_PATH": "/home/www/go/goAgile/", //安装目录(绝对路径)
}

//获取配置信息的
func Get(set string) string {
	return Config[set]
}

//gman数据库配置
func GetGmanConf() map[string]string{    
    gman_conf := make(map[string]string)
    gman_conf["dbhost"] = "tcp(9.9.9.9:3306)"
    gman_conf["dbuser"] = "root"
    gman_conf["dbpass"] = "123456"
    gman_conf["dbname"] = "gman_db"
    gman_conf["charset"] = "utf8"
    return gman_conf
}



