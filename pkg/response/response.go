/*
 * @Author: Go不浪队
 * @Date: 2023-02-05 20:38:51
 * @LastEditTime: 2023-02-07 00:08:23
 * @Description:
 */
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode	int         `json:"code"`
	StatusMsg	string      `json:"msg"`
	Data 		interface{} `json:",omitempty,inline"`
}

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  msg,
		Data: data,
	})
}

func Error(c *gin.Context, httpStatus int,errCode int, errMsg string) {
	c.JSON(httpStatus, Response{
		StatusCode: errCode,
		StatusMsg:  errMsg,
	})
	c.Abort()
}