package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
	"os"
	"io"
	"github.com/margin/server/db"
)

const PAGE_PATH = `D:\dev\server\go\src\github.com\margin\server\pages\`

func index(c *gin.Context)  {
	c.File(PAGE_PATH + `index.html`)
}

func favicon(c *gin.Context)  {
	c.File(PAGE_PATH + `res\favicon.ico`)
}

func form(c *gin.Context)  {
	if color, exist := c.GetPostForm("color"); exist {
		c.String(http.StatusOK, "color: %v\n", color)
	} else if datetime, exist := c.GetPostForm("datetime"); exist {
		c.String(http.StatusOK, "datetime: %v\n", datetime)
	} else if r, exist := c.GetPostForm("range"); exist {
		c.String(http.StatusOK, "range: %v\n", r)
	}

}

func datalist(c *gin.Context)  {
	if data, exist := c.GetPostForm("list"); exist {
		c.String(http.StatusOK, "列表值: %v\n", data)
	} else {
		c.String(http.StatusOK, "未选择列表")
	}
}

func file(c *gin.Context)  {
	filePart, err := c.FormFile("file")
	if err == nil {
		c.String(http.StatusOK,"name: %v, size: %d\n", filePart.Filename, filePart.Size)
		file, err := filePart.Open()
		if err != nil {
			c.String(http.StatusOK, "err: %v\n", err)
			return
		}
		defer file.Close()
		io.Copy(os.Stdout, file)
	} else {
		c.String(http.StatusNotFound, "err: %v\n", err)
	}
}

func indexV2(c *gin.Context)  {
	c.File(PAGE_PATH + `index_v2.html`)
}

func json(c *gin.Context)  {
	type Form struct {
		User string `form:"user"`
		Password string `form:"password"`
	}
	var form Form
	err := c.Bind(&form)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"user": form.User, 
			"password": form.Password,
		})
	} else {
		c.String(http.StatusNotFound, "err: %v\n", err)
	}
}

func user(c *gin.Context)  {
	c.File(PAGE_PATH + "user.html")
}

func addUser(c *gin.Context)  {
	var user db.User
	err := c.Bind(&user)
	if err == nil {
		ret, err := user.Add()
		if err != nil {
			c.String(http.StatusNotFound, "err: %v\n", err)
		} else {
			c.String(http.StatusOK, "ok, insert user at %d\n", ret)
		}
	} else {
		c.String(http.StatusNotFound, "err: %v\n", err)
	}
}

func userList(c *gin.Context)  {
	users, err := db.QueryAllUsers()
	if err != nil {
		c.String(http.StatusNotFound, "err: %v\n", err)
		return
	}
	c.HTML(http.StatusOK, "userlist.html", gin.H {
		"users": users, 
	})
}

func Run()  {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default();
	router.LoadHTMLGlob(PAGE_PATH + "templates/*")
	router.Static("/res", PAGE_PATH + "res")
	v1 := router.Group("/v1")
	{
		v1.GET("/", index)
		v1.GET("/favicon.ico", favicon)
		v1.POST("/form", form)
		v1.POST("/datalist", datalist)
		v1.POST("/file", file)
	}
	v2 := router.Group("/v2")
	{
		v2.GET("/", indexV2)
		v2.POST("/json", json)
	}
	v3 := router.Group("/v3")
	{
		v3.GET("/", user)
		v3.POST("/addUser", addUser)
		v3.GET("/userlist", userList)
	}	
	// router.GET("/:name", file)
	router.Run("192.168.1.105:3000")
}