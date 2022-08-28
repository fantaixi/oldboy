package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"oldboy/better/11grpc/service"
)

func main() {
	user := &service.User{
		Username: "aaa",
		Age: 18,
	}

	//序列化
	marshal, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}

	//反序列化
	newUser := &service.User{}
	err = proto.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(newUser.String())
}
