package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int64
	Name string
	Age  int
	Sex  int
}

func main() {
	//打开数据库
    //DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	db, err := sql.Open("mysql", "root:ozixue@/go_driver?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	//关闭数据库，db会被多个goroutine共享，可以不调用
	defer db.Close()

	stmtIns, err = db.Prepare("INSERT INTO user VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()
}
