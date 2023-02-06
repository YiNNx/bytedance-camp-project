/*
 * @Author: Go不浪队
 * @Date: 2023-02-05 20:21:15
 * @LastEditTime: 2023-02-07 02:08:10
 * @Description:
 */
package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"tiktok/cmd/api/controller/param"
	"tiktok/cmd/api/rpc"
	"tiktok/kitex_gen/user"
	"tiktok/pkg/response"
)

func Register(c *gin.Context) {
	var req param.ReqSignUp
	if err := c.ShouldBind(&req); err != nil {
		// TODO: 第三个参数为自定义错误码
		response.Error(c, http.StatusBadRequest, -1, err.Error())
		return
	}

	res,err:=rpc.Register(context.Background(), &user.ServiceReqRegister{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		response.Error(c, int(res.BaseRes.HttpCode), int(res.BaseRes.StatusCode), err.Error())
		return
	}
	if(res.BaseRes.StatusCode!=0){
		response.Error(c, int(res.BaseRes.HttpCode), int(res.BaseRes.StatusCode), res.BaseRes.StatusMessage)
		return
	} else {
		response.Success(c,res.BaseRes.StatusMessage,param.ResSignUp{
			UserId: res.UserId,
			Token:  res.Token,
		})
	}
}

func Login(c *gin.Context) {

}
