package thttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseFailDefault(c *gin.Context, msg string) {
	ResponseFail(c, http.StatusInternalServerError, http.StatusInternalServerError, msg, "")
}

func ResponseFail[T any](c *gin.Context, httpCode int, code int, msg string, data T) {
	c.JSON(httpCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
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
