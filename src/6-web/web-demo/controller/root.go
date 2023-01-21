package controller

import (
	"context"
	"fmt"
	"web-demo/model"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.String(200, "Hello World!")
}

func Post(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println("error!")
	}
	if HTTPCtx, ok := c.Get("ctx"); ok {
		if ctx, ok := HTTPCtx.(context.Context); ok {
			fmt.Println(ctx)
		}
	} else {
		fmt.Println(c.Request.Context())
	}
	c.JSON(200, gin.H{"id": user.Id, "age": user.Age})
}

func Route(r *gin.Engine) {
	r.GET("/", HelloWorld)
	r.GET("/hello", HelloWorld)
	r.GET("/hello/world", HelloWorld)
	r.POST("/", Post)
}
