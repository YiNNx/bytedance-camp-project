/*
 * @Author: Go不浪队
 * @Date: 2023-02-07 01:06:16
 * @LastEditTime: 2023-02-07 20:09:50
 * @Description:
 */
package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"

	user "tiktok/kitex_gen/user/userservice"
	"tiktok/pkg/config"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", config.C.Service.UserHostPort)
	svr := user.NewServer(new(UserServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
