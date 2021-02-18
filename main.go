package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//链接数据库
var db, _ = gorm.Open(mysql.Open(`root:123456@(127.0.0.1:3306)/rubbish?charset=utf8&parseTime=True&loc=Local`), &gorm.Config{PrepareStmt: true})

func main() {
	err := db.AutoMigrate(&User{}, &Recycle{}, &Rubbish{})
	if err != nil {
		fmt.Println(err)
	}
	gin.SetMode(gin.DebugMode) // 运行模式
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.POST("/login", login)
	router.POST("/register", register)
	router.Use(auth()) //鉴权
	router.Any("/recycle", recycle)
	router.GET("/search", search)

	err = router.Run(":8081")
	if err != nil {
		fmt.Println(err.Error())
	}
}
