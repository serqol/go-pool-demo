package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	amqp "github.com/serqol/go-amqp"
	redis "github.com/serqol/go-redis"
	utils "github.com/serqol/go-utils"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery())
	gin.SetMode(gin.DebugMode)

	router.POST("/", publishController)

	router.Run(utils.GetEnv("SOCKET", "localhost:8888").(string))
}

func publishController(c *gin.Context) {
	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)
	go publish(x)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"error":   0,
	})
}

func publish(body []byte) {
	_, error := amqp.GetDefaultConnector().Execute(func(connection interface{}) (interface{}, error) {
		publishError := connection.(*amqp.Amqp).Publish(body)
		return nil, publishError
	})
	if error != nil {
		fmt.Print(error)
	}
	_, redisError := redis.GenerateUniqueKey(string(body), time.Second*60, "go_pool_demo_")

	if redisError != nil {
		fmt.Print(redisError)
	}
}
