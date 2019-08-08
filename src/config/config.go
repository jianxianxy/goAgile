/**
  配置信息
  定义的配置信息使用小写字母，保持对外私有
*/
package config

import (
	"lib/db"
)

var Config = map[string]string{
	"ROOT_PATH": "/Project/Go/goAgile/",      //安装目录(绝对路径)
	"ROOT_VIEW": "/Project/Go/goAgile/view/", //视图目录
}

//获取配置信息的
func Get(set string) string {
	return Config[set]
}

//数据库链接（单例模式）
type dbStru struct {
	dbcsn db.Mysql
}

//gman数据库
var dbGman *dbStru

func DbGman() db.Mysql {
	if dbGman == nil {
		var dbc db.Mysql
		gman_conf := make(map[string]string)
		gman_conf["dbhost"] = "tcp(127.0.0.1:3306)"
		gman_conf["dbuser"] = "root"
		gman_conf["dbpass"] = "123"
		gman_conf["dbname"] = "gman_db"
		gman_conf["charset"] = "utf8"
		dbc.GetConn(gman_conf)
		dbGman = &dbStru{}
		dbGman.dbcsn = dbc
	}
	return dbGman.dbcsn
}

//spider数据库
var dbCspi *dbStru

func DbSpider() db.Mysql {
	if dbCspi == nil {
		var dbc db.Mysql
		spider_conf := make(map[string]string)
		spider_conf["dbhost"] = "tcp(127.0.0.1:3306)"
		spider_conf["dbuser"] = "root"
		spider_conf["dbpass"] = "123"
		spider_conf["dbname"] = "spider_db"
		spider_conf["charset"] = "utf8"
		dbc.GetConn(spider_conf)
		dbCspi = &dbStru{}
		dbCspi.dbcsn = dbc
	}
	return dbCspi.dbcsn
}
