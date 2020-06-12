package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
)

//获取请求的URL等信息
func handleClientRequest(client net.Conn) {

	if client == nil {
		return
	}

	defer client.Close()


	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	//var method, host, address string
	////字符串b截取第一个\n
	//fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)

	var urlString string = string(b[:bytes.IndexByte(b[:], '\r')])
	urlString = strings.TrimLeft(urlString, "GET ")
	urlString = strings.TrimRight(urlString, " HTTP/1.1")
	url, err := url.Parse(urlString)//String --> url
	if err != nil {
		log.Println(err)
		return
	}
	//解析该url
	var address string
	if url.Opaque == "443" {
		address = url.Scheme + ":443"
	} else {
		if strings.Index(url.Host, ":") == -1 {
			address = url.Host + ":80"
		} else {
			address = url.Host
		}
	}


	//根据该url的address，建立连接
	fmt.Println("收到转发请求， Target = ", address)
	server, err := net.Dial("tcp", address)
	if err != nil {
		 log.Println(err)
		 return
	}

	//数据转发
//	if method == "CONNECT" {
//		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	//} else {
		server.Write(b[:n])
	//}

	//client --> server 转发客户端信息
	go io.Copy(server, client)

	//server --> client 发回给客户端
	io.Copy(client, server)


}

func main() {

	//启动8080监听
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	//接收代理请求
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleClientRequest(client)
	}


}
