package toolkit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InternalError[T any](c *gin.Context, code int, msg string, data ...T) {
	ResponseFail(c, http.StatusInternalServerError, code, msg, data)
}
func ResponseFail[T any](c *gin.Context, httpCode int, code int, msg string, data ...T) {
	if len(data) > 0 {
		c.JSON(httpCode, gin.H{
			"code": code,
			"msg":  msg,
			"data": data[0],
		})
	} else {
		c.JSON(httpCode, gin.H{
			"code": code,
			"msg":  msg,
		})
	}
	return
}

func ResponseSuccess[T any](c *gin.Context, code int, msg string, data T) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
