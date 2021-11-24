package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func login(c *gin.Context) {
	p := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&p)
	var user User
	db.Where(`user = ? AND pwd = ? `, p["user"], p["pwd"]).Find(&user)
	if user.User != `` {
		session := strconv.Itoa(int(time.Now().UnixNano())) + p["user"].(string)
		db.Model(&user).Update(`session`, session)
		c.SetCookie("session", session, 100000, "/", "", false, false)
		c.SetCookie("type", user.Type, 100000, "/", "", false, false)
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
	p := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&p)
	fmt.Println(p["user"], p["pwd"])
	var user User
	db.Where(`user = ? `, p["user"]).Find(&user)
	if user.User != `` || p["user"] == `` {
		c.JSON(200, gin.H{
			"data": `用户名已存在`,
		})
		return
	}

	db.Create(&User{
		User:  p["user"].(string),
		Pwd:   p["pwd"].(string),
		Type:  p["type"].(string),
		Phone: p["phone"].(string),
	})
	c.JSON(200, gin.H{
		"data": `注册成功，请登录`,
	})

}

func recycle(c *gin.Context) {
	p := make(map[string]interface{}) //注意该结构接受的内容
	fmt.Println(c.Request.Method)
	switch c.Request.Method {
	case `GET`:
		var recycles []Recycle
		db.Find(&recycles)
		c.JSON(200, gin.H{
			"data": recycles,
		})
	case `PUT`:
		session, _ := c.Cookie(`session`)
		var user User
		db.Where(`user = ?`,session[19:]).Find(&user)
		db.Model(&Recycle{}).Where("id = ?", c.Query(`id`)).Updates(map[string]interface{}{
			"r_user":user.User,
			"r_phone":user.Phone,
		})
		c.JSON(200, gin.H{
			"data": "回收成功",
		})
	case `POST`:
		c.BindJSON(&p)
		session, _ := c.Cookie(`session`)
		db.Create(&Recycle{
			CUser: session[19:],
			Place: p["place"].(string),
			Time:  time.Now(),
			Info:  p["info"].(string),
		})
		c.JSON(200, gin.H{
			"data": "添加成功",
		})
	case `DELETE`:
		fmt.Println(c.Query(`id`))
		fmt.Println(db.Where("id = ?", c.Query(`id`)).Delete(&Recycle{}).Debug())
		c.JSON(200, gin.H{
			"data": "删除成功",
		})
	}

}

func search(c *gin.Context) {
	/*	resp, err := http.Get("https://api.66mz8.com/api/garbage.php?name=" + c.Query(`name`))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		var res map[string]interface{}
		_ = json.Unmarshal(body, &res)*/
	var g Garbage
	var g1 Garbage

	db.Where(`id = 1`).Find(&g1)
	fmt.Println(g1)

	db.Where(`name = ?`, c.Query(`name`)).Find(&g)

	c.JSON(200, gin.H{
		`data`: g,
	})
}
