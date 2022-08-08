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

//查询
/*
单行查询
单行查询db.QueryRow()执行一次查询，并期望返回最多一行结果（即Row）。
QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。（如：未找到结果）
func (db *DB) QueryRow(query string, args ...interface{}) *Row
*/
func queryRow(id int) {
	var u1 user
	//1、查询单条记录的SQL语句
	sqlStr := "select id, name, age from user where id=?;"
	//2、执行
	rowObj := Db.QueryRow(sqlStr,id)
	//3、拿到结果去映射
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	rowObj.Scan(&u1.id,&u1.name,&u1.age)
	/*
	// 或者2 3一起执行
		err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
	 */
	//打印
	fmt.Printf("%#v\n",&u1)
}

/*
多行查询
多行查询db.Query()执行一次查询，返回多行结果（即Rows），一般用于执行select命令。参数args表示query中的占位参数。
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
 */
func queryMore(id int) {
	sqlStr := "select id,name,age from user where id > ?"
	rows,err := Db.Query(sqlStr,id)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	//记得关闭连接
	defer rows.Close()
	//循环取值
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

/*
插入数据
插入、更新和删除操作都使用Exec方法。
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
Exec执行一次命令（包括查询、删除、更新、插入等），返回的Result是对已执行的SQL命令的总结。参数args表示query中的占位参数。
 */

func insert() {
	sqlStr := `insert into user(name,age) values("小小林",24)`
	ret,err := Db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	//如果是插入数据的操作，能够拿到插入数据的ID
	id,err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", id)
}

//更新数据
func updateRow(newAge int, id int) {
	//sql := `update user set age=? where id=?`
	//更新id大于?的数据
	sql := `update user set age=? where id>?`
	ret,err := Db.Exec(sql,newAge,id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func delete(id int) {
	sqlStr := "delete from user where id = ?"
	ret,err :=Db.Exec(sqlStr,id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initdb failed,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")
	//queryRow(1)
	//queryMore(0)
	//insert()
	//updateRow(99,1)
	delete(3)
}
