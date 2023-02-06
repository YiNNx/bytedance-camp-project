/*
 * @Author: Go不浪队
 * @Date: 2023-02-05 20:16:12
 * @LastEditTime: 2023-02-07 00:24:53
 * @Description:
 */
package router

import (
	"github.com/gin-gonic/gin"

	"tiktok/cmd/api/controller"
)

func initUserRouters(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}
}