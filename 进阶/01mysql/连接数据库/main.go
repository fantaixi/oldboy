package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //不使用，但是需要里面的init()函数
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/oldboy"
	db, err := sql.Open("mysql", dsn)  //不会校验用户名和密码是否正确
	if err != nil {
		fmt.Printf("dsn:%s failed,err:%v\n",dsn,err)
		return
	}
	defer db.Close()  // 注意这行代码要写在上面err判断的下面
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n",dsn,err)
		return
	}
	fmt.Println("连接数据库成功")
}
