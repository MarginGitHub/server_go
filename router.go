package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"os"
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
	}

}

func file(c *gin.Context)  {
	name := c.Param("name")
	fmt.Fprintf(os.Stdout, "%v\n", name)
	c.File(`C:\Users\Administrator\Desktop\wallpaper-1-sierra.jpg`)
}

func Run()  {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default();
	// router.StaticFile("/favicon.ico", ".\\favicon.ico")
	router.GET("/", index)
	router.GET("/favicon.ico", favicon)
	router.POST("/form", form)
	// router.GET("/:name", file)
	router.Run("127.0.0.1:3000")
}