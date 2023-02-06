/*
 * @Author: Go不浪队
 * @Date: 2023-02-05 20:38:51
 * @LastEditTime: 2023-02-06 20:59:35
 * @Description:
 */
package context

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
	Err  []string    `json:"error,omitempty"`
}

func Success(c *gin.Context, status int, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: status,
		Data: data,
		Msg:  msg,
	})
}

func Error(c *gin.Context, status int, msg string, err ...string) {
	c.JSON(status, Response{
		Code: status,
		Msg:  msg,
		Err:  err,
	})
}