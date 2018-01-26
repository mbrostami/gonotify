package routes

import (
	"github.com/gin-gonic/gin"
)

func define (router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/sms/single", smssingle)
		v1.POST("/sms/bulk", smsbulk)
	}
}