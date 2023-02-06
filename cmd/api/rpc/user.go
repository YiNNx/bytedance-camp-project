/*
 * @Author: Go不浪队
 * @Date: 2023-02-06 21:12:02
 * @LastEditTime: 2023-02-07 02:07:20
 * @Description:
 */
package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"

	"tiktok/kitex_gen/user"
	"tiktok/kitex_gen/user/userservice"
	"tiktok/pkg/config"
)


var userClient userservice.Client

func initUserRpc() {
	c, err := userservice.NewClient("example", client.WithHostPorts(config.UserHostPorts))
	if err != nil {
		panic(err)
	}
	userClient = c
}

func Register(ctx context.Context, req *user.ServiceReqRegister)(*user.ServiceResRegister,error) {
	return userClient.Register(ctx, req)
}

// TODO: 