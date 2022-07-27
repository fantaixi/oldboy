package main

import (
	"fmt"
	"os"
)

var (
	allStudent map[int64]*student
)

type student struct {
	id int64
	name string
}

func newStuent(id int64, name string) *student {
	return &student{
		id: id,
		name: name,
	}
}
func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号:%d  名字:%s\n",k,v.name)
	}
}
func addStudent() {
	var (
		id int64
		name string
	)
	fmt.Print("输入学号:")
	fmt.Scanln(&id)
	fmt.Print("输入名字:")
	fmt.Scanln(&name)
	newStu := newStuent(id,name)
	allStudent[id] = newStu
}
func deleteStudent() {
	var (
		deleteId int64
	)
	fmt.Print("输入要删除的学生学号:")
	fmt.Scanln(&deleteId)
	//直接用delete() 根据学号删除对应的键值对
	delete(allStudent,deleteId)
}
func main() {
	allStudent = make(map[int64]*student,20)
	for {
		//1、打印菜单
		fmt.Println("学生管理系统")
		fmt.Println(`
 		1、查看所有学生
		2、新增学生
		3、删除学生
		4、退出
	`)
		//2、等待用户选择
		fmt.Print("请选择：")
		var choose int
		fmt.Scanln(&choose)
		//3、执行对应的函数
		switch choose {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("滚")
		}
	}

}
