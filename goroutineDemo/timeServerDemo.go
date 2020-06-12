package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func handle(c net.Conn){
	defer c.Close()
	b:= make([]byte, 100)
	c.Read(b)
	for i := 0; i < 3;i ++ {
		_, err := io.WriteString(c, string(b))
		fmt.Println(string(b))
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		go handle(conn)
	}


}