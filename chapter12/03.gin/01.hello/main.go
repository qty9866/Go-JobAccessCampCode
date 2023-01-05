package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("你好,Gin!"))
	})
	err := r.Run(":8080")
	if err != nil {
		log.Println("運行錯誤", err)
	}
}
