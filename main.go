package main

import (
    // "fmt"
    // "reflect"
	"github.com/gin-gonic/gin"
	// "github.com/mbrostami/gonotify/routes"
)


func main() {
	router := gin.Default()
	// fmt.Println(reflect.TypeOf(router))
	v1 := router.Group("/v1")
	{
		v1.POST("/sms", sendsms)
		v1.GET("/gateways", gateways)
	}
	router.Run()
}

func sendsms(c *gin.Context) {
	id := c.Query("id")
	// page := c.DefaultQuery("page", "0")
	// name := c.PostForm("name")
	// message := c.PostForm("message")

	// fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	c.JSON(200, gin.H{
		"message": id,
	})
}
func gateways(c *gin.Context) {
	c.JSON(200, gin.H{
		"id": 12345,
		"name": "gateway name",
		"bulkSupport": true,
	})
}