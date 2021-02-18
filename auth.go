package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		session, _ := c.Cookie(`session`)

		fmt.Println(`session`, session)
		db.Where(`session = ?`, session).Find(&user)
		if user.User != `` && user.Session != `` {
			c.Next()
		} else {
			c.JSON(200, gin.H{
				`data`: `session验证失败`,
			})
			c.Abort()
		}
	}
}
