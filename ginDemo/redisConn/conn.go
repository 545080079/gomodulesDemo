package redisConn

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"os"
	"strings"
)


func GetConn() redis.Conn{
	file, err := os.Open("redis.plain")
	b := make([]byte,100)
	file.Read(b)
	login := strings.Split(string(b), "\n")

	conn, err := redis.Dial("tcp",login[0])
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Do("auth", login[1])
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
