package main

import (
	"demo/ginDemo/redisConn"
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)


func main() {

	//连接redis
	redisConn := redisConn.GetConn()
	defer redisConn.Close()
	//设置http服务

	req := gin.Default()

	ginpprof.Wrap(req)
	req.GET("/a", func (c *gin.Context) {
		c.IndentedJSON(200, gin.H{	//gin.H : 一个map[string]interface{}对象
			"Type" : "GET",
			"Name" : "lyt",
			"WeChat" : "545080079",
		})
	})
	req.POST("/a", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"Type" : "POST",
			"Name" : "lyt",
			"WeChat" : "545080079",
		})
	})

	//路由参数
	req.GET("/users/:id", func (c *gin.Context) {
		id := c.Param("id")
		c.String(200, "用户%s登陆成功: paramX = " + c.Query("X"), id)
	})
	req.GET("/users", func (c *gin.Context) {
		c.String(200, "Default Admin Login OK")
	})


	v1 := req.Group("/v1")

	//由于中间件是递归调用，所以先调用最先加入数组的中间件（先进先出）
	v1.Use(middlewareLogin)
	v1.Use(middlewarePre)
	v1.GET("/:id", func (c *gin.Context) {
		id := c.Param("id")
		val, ok := c.Get("auth")
		if ok == false {
			c.String(200, "登陆失败! 未知错误\n")
		}else {
			if val == false {
				c.String(200, "登陆失败! 账号认证失败\n")
			}else {
				c.String(200, "用户%s登陆成功: paramX = " + c.Query("X") + "\n", id)
			}
		}
	})



	apiQ := req.Group("/q")
	apiQ.GET("/:key", func (c *gin.Context) {
		key := c.Param("key")
		res, err := redis.String(redisConn.Do("GET", key))
		if err != nil {
			log.Fatal(err)
		}
		c.String(200, "找到数据：")
		c.IndentedJSON(200, gin.H{
			key : res,
		})

	})


	req.Run(":8080")
}

func middlewarePre(c *gin.Context) {
	fmt.Println("Welcome !")
}

func middlewareLogin(c *gin.Context){
	fmt.Println("调用middleware：Login")
	if !strings.EqualFold(c.Param("id"), "123") {
		fmt.Println("认证失败！")
		c.Set("auth", false)
	}else {
		fmt.Println("认证成功。")
		c.Set("auth", true)
	}
}