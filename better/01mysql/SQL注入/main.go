package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //不使用，但是需要里面的init()函数
)

var Db *sql.DB  //底层是一个数据库连接池

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/oldboy"
	Db, err = sql.Open("mysql", dsn) //不会校验用户名和密码是否正确
	if err != nil {
		fmt.Printf("dsn:%s failed,err:%v\n", dsn, err)
		return
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = Db.Ping()
	if err != nil {
		return
	}
	//设置与数据库建立连接的最大数目
	Db.SetMaxOpenConns(10)
	//设置连接池中的最大闲置连接数
	Db.SetMaxIdleConns(2)
	return
}

type user struct {
	id   int
	name string
	age  int
}

// sql注入示例
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u user
	err := Db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initdb failed,err:%v\n", err)
		return
	}
	sqlInjectDemo("xxx' or 1=1#")
	sqlInjectDemo("xxx' union select * from user #")
	sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
}
