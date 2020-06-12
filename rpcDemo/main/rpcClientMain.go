package main

import (
	"context"
	rpcClient "demo/rpcDemo/rpcServer"
	"google.golang.org/grpc"
	"io"
	"log"
)

const ServerAddr = "127.0.0.1:9090"
func main() {
	conn, err := grpc.Dial(ServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	//新建rpc客户端
	rpcClientInstance := rpcClient.NewHelloClient(conn)

	res, err := rpcClientInstance.SayHello(context.Background(), &rpcClient.HelloRequest{
		Name: "Hello RPC!",
	})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("res = ", res.Message)

	stream, err := rpcClientInstance.LotsOfReplies(context.Background(), &rpcClient.HelloRequest{
		Name: "Hello RPC!",
	})

	if err != nil {
		log.Fatalln(err)
	}

	for {
		res , err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%s", res.Message)
	}

}
