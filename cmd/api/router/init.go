/*
 * @Author: Go不浪队
 * @Date: 2023-02-05 20:15:56
 * @LastEditTime: 2023-02-07 02:34:02
 * @Description:
 */

package router

import (
	"github.com/gin-gonic/gin"

	"tiktok/cmd/api/controller"
)

func InitRouters(r *gin.RouterGroup) {
	initUserRouters(r)
}

func initUserRouters(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}
}

// TODO