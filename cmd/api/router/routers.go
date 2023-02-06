/*
 * @Author: Go不浪队
 * @Date: 2023-02-05 20:15:56
 * @LastEditTime: 2023-02-06 20:15:56
 * @Description:
 */

package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	initUserRouters(r)
}