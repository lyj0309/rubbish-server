package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func login(c *gin.Context) {
	var user User
	db.Where(`user = ? AND pwd = ? `, c.PostForm(`user`), c.PostForm(`pwd`)).Find(&user)
	if user.User != `` {
		session := strconv.Itoa(int(time.Now().UnixNano())) + c.PostForm(`user`)
		db.Model(&user).Update(`session`, session)
		c.SetCookie("session", session, 100000, "/", "", false, true)
		c.JSON(200, gin.H{
			"data": `登录成功`,
		})
	} else {
		c.JSON(200, gin.H{
			"data": `用户名或密码错误`,
		})
	}

}

func register(c *gin.Context) {
	var user User
	db.Where(`user = ? AND pwd = ? `, c.PostForm(`user`), c.PostForm(`pwd`)).Find(&user)
	if user.User != `` || c.PostForm(`user`) == `` {
		c.JSON(200, gin.H{
			"data": `用户名已存在`,
		})
		return
	}

	db.Create(&User{
		User: c.PostForm(`user`),
		Pwd:  c.PostForm(`pwd`),
	})
	c.JSON(200, gin.H{
		"data": `注册成功，请登录`,
	})

}

func recycle(c *gin.Context) {
	switch c.Request.Method {
	case `GET`:
		var recycles []Recycle
		db.Find(&recycles)
		c.JSON(200, gin.H{
			"data": recycles,
		})
	case `PUT`:
		session, _ := c.Cookie(`session`)
		db.Model(&Recycle{}).Where("id = ?", c.Query(`id`)).Update("r_user", session[19:])
		c.JSON(200, gin.H{
			"data": "选择成功",
		})
	case `POST`:
		session, _ := c.Cookie(`session`)
		db.Create(&Recycle{
			CUser: session[19:],
			Place: c.PostForm(`place`),
			Time:  time.Now(),
			Info:  c.PostForm(`info`),
		})
		c.JSON(200, gin.H{
			"data": "添加成功",
		})
	case `DELETE`:
		db.Where("id = ?", c.Query(`id`)).Delete(&Recycle{})
	}

}

func search(c *gin.Context) {
	resp, err := http.Get("https://api.66mz8.com/api/garbage.php?name=" + c.Query(`name`))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	var res map[string]interface{}
	_ = json.Unmarshal(body, &res)

	c.JSON(200, gin.H{
		`data`: res["data"],
	})
}
