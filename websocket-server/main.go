package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"syscall"

	"github.com/JoTaeYang/go-cook/common/cache"

	"github.com/JoTaeYang/go-cook/websocket-server/api"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func InitConfig() {
	// dbConfig := dynamodb.Config{
	// 	Region:    "",
	// 	AccessKey: "",
	// 	SecretKey: "",
	// 	Table:     "",
	// }

	//dynamodb.InitDynamoDB(dbConfig)
	// 	Region:    "ap-southeast-1",
	// 	AccessKey: "AKIAVU5C7TBCVMDCZ3T7",
	// 	SecretKey: "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",
	// 	Table:     "battle-league",
	// }

	// cache.InitRedis(&cache.Config{
	// 	[]string{
	// 		"192.168.2.240:6380",
	// 		"192.168.2.240:6381",
	// 		"192.168.2.240:6382",
	// 	},
	// }, true)
}

func InitGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//r.Use(SessionCheck())

	r.GET("/", api.Upgrader)
	return r
}

func main() {
	// Increase resources limitations
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	// Enable pprof hooks
	go func() {
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Fatalf("pprof failed: %v", err)
		}
	}()

	// Start epoll
	var err error
	err = api.MkEpoll()
	if err != nil {
		panic(err)
	}

	go api.StartEpoll()

	InitConfig()
	engine = InitGin()
	engine.Run(":50000")
}

func SessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			log.Println("not recv auth token")
			c.JSON(http.StatusOK, gin.H{
				"err_code": 1000,
			})
			return
		}

		log.Println("user connect session key : ", token)

		uid, err := cache.GetUid(token)
		if err != nil || uid == "" {
			log.Println("not found uid")
			c.JSON(http.StatusOK, gin.H{
				"err_code": 1000,
			})
			return
		}

		c.Request.Header.Set("UID", uid)
		c.Next()
	}
}
