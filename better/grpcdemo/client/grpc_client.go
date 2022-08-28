package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpcdemo/client/auth"
	"grpcdemo/service"
	"io/ioutil"
	"log"
)

func main() {
	//使用证书
	//creds,err := credentials.NewClientTLSFromFile("cert/server.pem","*.fantaixi.com")
	//if err != nil {
	//	log.Fatal("证书错误",err)
	//}

	// 证书认证-双向认证
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.crt")
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ServerName: "*.fantaixi.com",
		RootCAs:    certPool,
	})

	//token
	token := &auth.Authentication{
		User: "admin",
		Password: "admin",
	}
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(creds),grpc.WithPerRPCCredentials(token))
	if err != nil {
		log.Fatal("服务端出错，",err)
	}
	defer conn.Close()
	prodClient := service.NewProdServiceClient(conn)
	request := &service.ProductRequest{
		ProdId: 123,
	}
	stockResponse, err := prodClient.GetProductStock(context.Background(), request)
	if err != nil {
		log.Fatal("查询库存出错，",err)
	}
	fmt.Println("查询成功",stockResponse.ProdStock)
}
