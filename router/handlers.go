package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hi from gin")
}
func collectHandler(c *gin.Context) {
	category := c.Param("category")
	news.Collect(category)
	c.String(http.StatusOK, "Search is initialized")
}
func resultHandler(c *gin.Context) {
	category := c.Param("category")
	topics := news.Result(category)
	c.JSON(http.StatusOK, topics)
}
