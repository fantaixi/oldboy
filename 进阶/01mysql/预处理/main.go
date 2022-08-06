package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //不使用，但是需要里面的init()函数
)

/*
什么是预处理？

普通SQL语句执行过程：
客户端对SQL语句进行占位符替换得到完整的SQL语句。
客户端发送完整SQL语句到MySQL服务端
MySQL服务端执行完整的SQL语句并将结果返回给客户端。

预处理执行过程：
把SQL语句分成两部分，命令部分与数据部分。
先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
MySQL服务端执行完整的SQL语句并将结果返回给客户端。

为什么要预处理？
优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
避免SQL注入问题。

Go实现MySQL预处理
database/sql中使用下面的Prepare方法来实现预处理操作。
func (db *DB) Prepare(query string) (*Stmt, error)
Prepare方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。
 */

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

//预处理方式插入多条数据
func prepareInsert() {
	sqlStr := "insert into user(name,age) values(?,?)"
	stmt,err := Db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	//后续只需要拿到stmt去执行一些操作
	var m = map[string]int{
		"小小林":1,
		"大大林":88,
		"大小林":32,
	}
	for k, v := range m {
		stmt.Exec(k,v)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initdb failed,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")
	prepareInsert()
}
