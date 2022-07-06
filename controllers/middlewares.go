package controllers

import (
	"github.com/gin-gonic/gin"
)

func SetHeaders(c *gin.Context) {
	origin := c.Request.Header.Get("origin")

	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, origin")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
}
