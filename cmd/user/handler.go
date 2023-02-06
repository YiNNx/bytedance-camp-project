/*
 * @Author: Go不浪队
 * @Date: 2023-02-07 01:06:16
 * @LastEditTime: 2023-02-07 02:04:01
 * @Description:
 */
package main

import (
	"context"
	"net/http"

	"tiktok/cmd/user/command"
	"tiktok/db"
	user "tiktok/kitex_gen/user"
	"tiktok/pkg/bcrypt"
	"tiktok/pkg/jwt"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
// TODO: 错误码
func (s *UserServiceImpl) Register(ctx context.Context, req *user.ServiceReqRegister) (resp *user.ServiceResRegister, err error) {
	c:=db.GetContext()
	defer c.Close()

	//检查用户名是否重复
	duplicate, err := command.UsernameDuplicate(c,req.Username)
	if err != nil {
		return &user.ServiceResRegister{
			BaseRes: &user.ServiceResBasic{HttpCode: http.StatusInternalServerError, StatusCode: -1},
		},err 
	}
	if duplicate {
		return &user.ServiceResRegister{
			BaseRes: &user.ServiceResBasic{HttpCode: http.StatusInternalServerError, StatusCode: -1},
		},err 
	}
	
	// 加密
	var passwordHash []byte
	if passwordHash, err = bcrypt.HashPassword(req.Password); err != nil {
		return &user.ServiceResRegister{
			BaseRes: &user.ServiceResBasic{HttpCode: http.StatusInternalServerError, StatusCode: -1},
		},err 
	}

	//插入记录
	userID,err := command.CreateUser(c,
		&db.User{
			PasswordHash:  string(passwordHash),
			Username:      req.Username,
		},
	)
	if err != nil {
		c.Abort()
		return &user.ServiceResRegister{
			BaseRes: &user.ServiceResBasic{HttpCode: http.StatusInternalServerError, StatusCode: -1},
		},err 
	}

	//合成token
	token, err := jwt.GenerateToken(int(userID))
	if err != nil {
		c.Abort()
		return &user.ServiceResRegister{
			BaseRes: &user.ServiceResBasic{HttpCode: http.StatusInternalServerError, StatusCode: -1},
		},err 
	}

	if err := c.Commit(); err != nil {
		return &user.ServiceResRegister{
			BaseRes: &user.ServiceResBasic{HttpCode: http.StatusInternalServerError, StatusCode: -1},
		},err 
	}

	// 成功返回
	return &user.ServiceResRegister{
		BaseRes: &user.ServiceResBasic{
			StatusMessage: "register successfully",
		},
		UserId:  int64(userID),
		Token:   token,
	},nil 
}

// Signin implements the UserServiceImpl interface.
func (s *UserServiceImpl) Signin(ctx context.Context, serviceReqSignin *user.ServiceReqSignin) (resp *user.ServiceResSignin, err error) {
	// TODO: Your code here...
	return
}

// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, serviceReqGetUserById *user.ServiceReqGetUserById) (resp *user.ServiceResGetUserById, err error) {
	// TODO: Your code here...
	return
}
