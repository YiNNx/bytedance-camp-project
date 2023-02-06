/*
 * @Author: Go不浪队
 * @Date: 2023-02-05 20:16:12
 * @LastEditTime: 2023-02-06 20:24:22
 * @Description:
 */
package router

import (
	"github.com/gin-gonic/gin"

	"tiktok/cmd/api/controller"
)

func initUserRouters(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/signup", controller.SignUp)
		user.POST("/signin", controller.SignIn)
	}
}