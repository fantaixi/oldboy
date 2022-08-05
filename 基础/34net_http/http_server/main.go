package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	b,err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		fmt.Println(err)
	}
	w.Write(b)
}
func f2(w http.ResponseWriter, r *http.Request) {
	//拿到get请求中的参数
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name,age)

	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}
func main() {
	http.HandleFunc("/",f1)
	http.HandleFunc("/xxx",f2)
	http.ListenAndServe("127.0.0.1:9090",nil)
}
