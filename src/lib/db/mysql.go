package db

import (
	"strings"
	"database/sql"
	_ "lib/db/github.com/go-sql-driver/mysql"
)

type Mysql struct{
    db *sql.DB
}


//创建数据库连接
func (conn *Mysql) GetConn(info map[string]string) {
	db, err := sql.Open("mysql",info["dbuser"]+":"+info["dbpass"]+"@"+info["dbhost"]+"/"+info["dbname"]+"?charset="+info["charset"])
	checkErr(err)
    conn.db = db
}

//插入
func (conn *Mysql) Insert(table string,data map[string]interface{}) int{
    var col,_col string
    var din []interface{}
    for key,val := range data{
        col += key+","
        _col += "?,"
        din = append(din,val)    
    }
	stmt, err := conn.db.Prepare(`INSERT user (` + strings.Trim(col,",") +`) values (` + strings.Trim(_col,",") + `)`)
	checkErr(err)
	res, err := stmt.Exec(din...)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	return int(id)
}

//查询
func (conn *Mysql) GetRow(sql string) []map[string]string{
	rows, err := conn.db.Query(sql)
	checkErr(err)
	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
    
    //结果集
    var data []map[string]string
    data = make([]map[string]string,0)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
        data = append(data,record)
	}
    return data
}

//更新数据
func (conn *Mysql) Update(table string,data map[string]interface{},wdata map[string]interface{}) int{
    var col,where string
    var cval []interface{}
    for key,val := range data{
        col += key+"=?,"
        cval = append(cval,val)
    }
    for key,val := range wdata{
        where += " " + key + "=? AND"
        cval = append(cval,val)
    }
	stmt, err := conn.db.Prepare(`UPDATE ` + table + ` SET ` + strings.Trim(col,",") + ` WHERE ` + strings.Trim(where,"AND"))
	checkErr(err)
	res, err := stmt.Exec(cval...)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	return int(num)
}

//删除数据
func (conn *Mysql) Remove(table string,wdata map[string]interface{}) int{
    var where string
    var cval []interface{}
    for key,val := range wdata{
        where += " " + key + "=? AND"
        cval = append(cval,val)
    }
	stmt, err := conn.db.Prepare(`DELETE FROM ` + table + ` WHERE ` + strings.Trim(where,"AND"))
	checkErr(err)
	res, err := stmt.Exec(cval...)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	return int(num)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}