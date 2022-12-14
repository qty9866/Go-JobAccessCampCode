package main

import (
	"Learning-JobAccess-Camp/chapter12/02.practice/frinterface"
	"Learning-JobAccess-Camp/pkg/apis"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func connectDB() *gorm.DB {
	connection, err := gorm.Open(mysql.Open("root:199866@tcp(127.0.0.1:3306)/go_db"))
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	fmt.Println("数据库连接成功")
	return connection
}

func main() {
	conn := connectDB()
	var rankServer = NewDbRank(conn, NewFatRateRank())
	// 类型断言是否成功
	if initRank, ok := rankServer.(frinterface.RankInitInterface); ok {
		// 类型断言成功 开始初始化
		if err := initRank.Init(); err != nil {
			log.Fatal("初始化失败", err)
		}
	}

	r := gin.Default()
	pprof.Register(r)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello,this is index of gin",
		})
	})
	r.POST("/register", func(c *gin.Context) {
		var pi *apis.PersonalInformation
		err := c.BindJSON(&pi)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"ErrMessage": "无法解析注册信息",
			})
			return
		}
		if err = rankServer.RegisterPersonalInformation(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"ErrMessage": "注册失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Message": "Succeed",
		})
	})

	// 使用PUT进行更新操作
	r.PUT("/personalinfo", func(c *gin.Context) {
		var pi *apis.PersonalInformation
		err := c.BindJSON(&pi)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"ErrMessage": "无法解析注册信息",
			})
			return
		}
		if resp, err := rankServer.UpdatePersonalInformation(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"ErrMessage": "注册失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, resp)
		}
	})

	r.GET("/rank/:name", func(c *gin.Context) {
		name := c.Param("name")
		fr, err := rankServer.GetFatRate(name)
		if err != nil {
			c.JSON(400, gin.H{
				"ErrMessage": "获取个人排行失败",
			})
			return
		} else {
			c.JSON(200, fr)
		}
	})

	r.GET("/ranktop", func(c *gin.Context) {
		top, err := rankServer.GetTop()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"ErrMessage": "获取排行榜失败",
			})
		} else {
			c.JSON(200, top)
		}

	})
	err := r.Run(":8081")
	if err != nil {
		//todo handle error
	}
}
