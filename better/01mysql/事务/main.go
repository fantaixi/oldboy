package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //不使用，但是需要里面的init()函数
)

/*
什么是事务？
事务：一个最小的不可再分的工作单元；通常一个事务对应一个完整的业务(例如银行账户转账业务，该业务就是一个最小的工作单元)，
同时这个完整的业务需要执行多次的DML(insert、update、delete)语句共同联合完成。A转账给B，这里面就需要执行两次update操作。

在MySQL中只有使用了Innodb数据库引擎的数据库或表才支持事务。
事务处理可以用来维护数据库的完整性，保证成批的SQL语句要么全部执行，要么全部不执行。

事务的ACID
通常事务必须满足4个条件（ACID）：
原子性（Atomicity，或称不可分割性）、一致性（Consistency）、隔离性（Isolation，又称独立性）、持久性（Durability）。

条件	解释
原子性	一个事务（transaction）中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。
		事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。
一致性	在事务开始之前和事务结束以后，数据库的完整性没有被破坏。
		这表示写入的资料必须完全符合所有的预设规则，这包含资料的精确度、串联性以及后续数据库可以自发性地完成预定的工作。
隔离性	数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。
		事务隔离分为不同级别，包括读未提交（Read uncommitted）、读提交（read committed）、可重复读（repeatable read）和串行化（Serializable）。
持久性	事务处理结束后，对数据的修改就是永久的，即便系统故障也不会丢失。
事务相关方法
Go语言中使用以下三个方法实现MySQL中的事务操作。 开始事务

func (db *DB) Begin() (*Tx, error)
提交事务

func (tx *Tx) Commit() error
回滚事务

func (tx *Tx) Rollback() error
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

func transactionDemo() {
	//开启事务
	tx,err := Db.Begin()
	if err != nil {
		fmt.Printf("start transaction failed,err:%v",err)
	}
	//执行多个SQL操作
	sqlStr1 := "update user set age=age-2 where id=1"
	sqlStr2 := "update xxx set age=age+2 where id=2"
	//2、执行
	_,err = tx.Exec(sqlStr1)
	if err != nil {
		//失败就回滚
		tx.Rollback()
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	_,err = tx.Exec(sqlStr2)
	if err != nil {
		//失败就回滚
		tx.Rollback()
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	//上面两个SQL都执行成功，就提交事务
	err = tx.Commit()
	if err != nil {
		//失败就回滚
		tx.Rollback()
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}
	fmt.Println("事务执行成功!")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initdb failed,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")
	transactionDemo()
}

