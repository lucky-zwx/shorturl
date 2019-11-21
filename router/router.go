package router

import (
	"awesomeProject/zwxurl/uid"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/wonderivan/logger"
	"net/http"
	"time"
)

var router = gin.Default()
var re_client, err = RedisInit("127.0.0.1:6379")

func RedisInit(addr string) (*redis.Client, error) {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "abk6745741",
		DB:       0,
		PoolSize: 10,
	})
	_, err := redisdb.Ping().Result()
	if err != nil {
		return nil, err
	} else {
		return redisdb, nil
	}
}

func init() {
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	router.LoadHTMLGlob("templates/*")
	router.Use(cors.Default())
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	index()
	if err := router.Run(":80"); err != nil {
		logger.Error("服务器启动失败")
	}
	defer re_client.Close()
}

func index() {
	router.GET("/-/:shorturl", func(c *gin.Context) {
		shorturl := c.Param("shorturl")
		ret, err := re_client.HGet("h1", shorturl).Result()
		if err != nil {
			fmt.Println(err)
		}
		c.Redirect(http.StatusMovedPermanently, ret)
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/", func(c *gin.Context) {
		key := uid.DecimalToAny(int(time.Now().UnixNano()), 76)
		fmt.Println(key)
		message := c.PostForm("url")
		_, err := re_client.HSet("h1",key, message).Result()
		if err != nil {
			fmt.Println(err)
		}
		c.String(http.StatusOK, "http://zwxurl.top/-/"+key)
	})
}
