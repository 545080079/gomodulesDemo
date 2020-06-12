package main

import (
	"demo/rpcDemo/rpcServer"
	"google.golang.org/grpc"
	"log"
	"net"
)
const(
	Address = "127.0.0.1:9090"
)
func main() {

	listener, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	rpcServer.RegisterHelloServer(s, &rpcServer.ApiController{})
	log.Println("Listen on " , Address)

	if err := s.Serve(listener); err != nil  {
		log.Fatalln("Fail to serve : ", err)
	}

}