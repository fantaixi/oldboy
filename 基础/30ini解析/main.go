package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
}

func loadIni(fileName string, data interface{}) (err error) {
	//1、参数校验
	//1.1 传进来的data必须是指针类型（需要在函数中对其赋值）
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer.")
		return
	}
	//1.2 data参数必须是结构体指针（配置文件中各种键值对需要赋值给对应的结构体）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data should be a struct.")
		return
	}
	//2、读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//将字节类型的文件内容转换成字符串并且切割
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Printf("%#v\n", lineSlice)
	//3、一行一行的读
	var structName string
	for idx, line := range lineSlice {
		//去掉字符串的首位空格
		line = strings.TrimSpace(line)
		//有空行就跳过
		if len(line) == 0 {
			continue
		}
		//3.1 如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//3.2 如果是[]开头就表示是节
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			//如果是  [  ]  的情况
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			//根据sectionName去data里面根据反射找到对应的结构体
			//v := reflect.ValueOf(data)
			for i := 0; i < t.Elem().NumField(); i++ {
				filed := t.Elem().Field(i)
				if sectionName == filed.Tag.Get("ini") {
					//说明找到对应的嵌套结构体，把字段名记录下来
					structName = filed.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			//3.3 如果不是[]开头的就是要找的以=分割的键值对
			//1、以等号分割这一行，等号左边是key，右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//2、根据structName去data里面把对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) //拿到嵌套结构体的值信息
			sType := sValue.Type()                     //拿到嵌套结构体的类型信息
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的:%s字段应该是一个结构体\n", structName)
				return
			}
			var fliedName string
			var fileType reflect.StructField
			//3、遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i) //tag信息是存储在类型信息中的
				fileType = filed
				if filed.Tag.Get("ini") == key {
					//找到对应的字段
					fliedName = filed.Name
					break
				}
			}
			//4、如果key=tag，给这个字段赋值
			if len(fileName) == 0 {
				//在结构体中找不到对应的字段
				continue
			}
			// 根据fieldName去取出这个字段
			fileObj := sValue.FieldByName(fliedName)
			//对其赋值
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool,err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d type error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			}
		}
	}
	return
}
func main() {
	var mc Config
	err := loadIni("./conf.ini", &mc)
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
	}
	fmt.Println(mc)
}
