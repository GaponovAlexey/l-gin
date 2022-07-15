package router

import (
	"github.com/gin-gonic/gin"
	"l/gin/news"
	"net/http"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hi from gin")
}
func collectHandler(c *gin.Context) {
	category := c.Param("category")
	news.Collect(category)
	c.String(http.StatusOK, "Search is initialized", category)
}
func resultHandler(c *gin.Context) {
	category := c.Param("category")
	topics := news.Result(category)
	c.JSON(http.StatusOK, topics)
}
