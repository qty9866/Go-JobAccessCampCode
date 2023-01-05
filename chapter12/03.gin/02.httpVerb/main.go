package main

import (
	"Learning-JobAccess-Camp/pkg/apis"
	"encoding/base64"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 一组请求 为了调试 / Debug 就可以加这个
	pprof.Register(r)
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte("This is index page from Gin!Welcome"))
	})
	// 如果用JSON 就可以返回JSON格式，并自带content-type: application/json
	r.GET("/history", func(c *gin.Context) {
		c.JSON(200, []apis.PersonalInformation{
			{
				Name:   "Qty",
				Sex:    "男",
				Tall:   1.81,
				Weight: 81.0,
				Age:    24,
			},
			{
				Name:   "Monica",
				Sex:    "女",
				Tall:   1.65,
				Weight: 47.0,
				Age:    22,
			},
		})
	})
	// 一般来说使用POST来进行创建
	r.POST("/register", func(c *gin.Context) {
		pi := &apis.PersonalInformation{}
		err := c.BindJSON(pi)
		if err != nil {
			//H is a shortcut for map[string]interface{}
			c.JSON(400, gin.H{
				"message": "无法读取个人注册信息",
			})
			return
		}
		//todo register to rank
		c.JSON(200, nil)
	})

	r.GET("/getWithQuery", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(200, gin.H{
			"payload": base64.StdEncoding.EncodeToString([]byte(name)),
		})
	})
	r.GET("/getWithParam/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"payload": base64.StdEncoding.EncodeToString([]byte(name)),
		})
	})
	err := r.Run(":8081")
	if err != nil {
		//todo handle err
	}
}
