/*
 * @Author: Go不浪队
 * @Date: 2023-02-05 19:31:45
 * @LastEditTime: 2023-02-06 20:34:29
 * @Description:
 */

package api

import (
	"github.com/gin-gonic/gin"

	"tiktok/cmd/api/router"
	"tiktok/pkg/config"
)

func main() {
	r := gin.Default()
	router.InitRouters(r)
	r.Run(config.C.App.Addr)
}