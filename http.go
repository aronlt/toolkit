package toolkit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseFailDefault[T any](c *gin.Context, msg string, data ...T) {
	ResponseFail(c, http.StatusInternalServerError, http.StatusInternalServerError, msg, data)
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

func ResponseSuccessDefault[T any](c *gin.Context, data T) {
	ResponseSuccess(c, 0, "success", data)
}

func ResponseSuccess[T any](c *gin.Context, code int, msg string, data T) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
