package main

import (
	"context"
	user "tiktok/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// ContainsName implements the UserServiceImpl interface.
func (s *UserServiceImpl) ContainsName(ctx context.Context, containsNameReq *user.ContainsNameReq) (resp *user.ContainsNameResp, err error) {
	// TODO: Your code here...
	return
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, createUserReq *user.CreateUserReq) (resp *user.CreateUserResp, err error) {
	// TODO: Your code here...
	return
}

// GetUserByName implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserByName(ctx context.Context, getUserByNameReq *user.GetUserByNameReq) (resp *user.GetUserByNameResp, err error) {
	// TODO: Your code here...
	return
}

// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, getUserByIdReq *user.GetUserByIdReq) (resp *user.GetUserByIdResp, err error) {
	// TODO: Your code here...
	return
}
