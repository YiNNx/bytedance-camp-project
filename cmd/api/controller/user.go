/*
 * @Author: Go不浪队
 * @Date: 2023-02-05 20:21:15
 * @LastEditTime: 2023-02-06 20:27:45
 * @Description:
 */
package controller

import (
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	// var err error
	// var json param.ReqSignUp
	// if err := c.ShouldBindJSON(&json); err != nil {
	// 	response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
	// 	return
	// }
	// hashedPassword, err := util.HashPassword(json.Password)
	// if err != nil {
	// 	response.Error(c, http.StatusInternalServerError, "Failed to hash password! ", err.Error())
	// 	return
	// }
	// m := model.GetModel(&model.UserModel{})

	// _, err = m.(*model.UserModel).FindUserByUsername(json.Username)
	// if err == nil {
	// 	response.Error(c, http.StatusBadRequest, "Username already exists. ")
	// 	return
	// }

	// u := &model.User{
	// 	Username:         json.Username,
	// 	Email:            json.Email,
	// 	Password:         hashedPassword,
	// 	Intro:            json.Intro,
	// 	Github:           json.Github,
	// 	School:           json.School,
	// 	Website:          json.Website,
	// 	UpdatePasswordAt: time.Now(),
	// }
	// id, err := m.(*model.UserModel).CreateUser(u)
	// if err != nil {
	// 	response.Error(c, http.StatusInternalServerError, "Failed to create user! ", err.Error())
	// 	return
	// }
	// response.Success(c, http.StatusOK, gin.H{"id": id, "role": "common"}, "Sign up successfully! ")
}

func SignIn(c *gin.Context) {
	// var json param.ReqSignIn
	// if err := c.ShouldBindJSON(&json); err != nil {
	// 	response.Error(c, http.StatusBadRequest, "Failed to bind json! ", err.Error())
	// 	return
	// }
	// m := model.GetModel(&model.UserModel{})
	// u, err := m.(*model.UserModel).FindUserByUsername(json.Username)
	// if err != nil {
	// 	response.Error(c, http.StatusBadRequest, "User does not exist. ", err.Error())
	// }

	// if ok, err := util.CheckPasswordHash(json.Password, u.Password); !ok {
	// 	response.Error(c, http.StatusBadRequest, "Password is wrong! ", err.Error())
	// 	return
	// }

	// //token, err := jwt.GetToken(json.Username, u.Role)
	// token := jwt.GenerateUserJwt(json.Username, u.Role, strconv.Itoa(int(u.ID)))
	// tokenStr, err := token.GenerateTokenStr()
	// if err != nil {
	// 	response.Error(c, http.StatusInternalServerError, "Failed to get token! ", err.Error())
	// 	return
	// }
	// response.Success(c, http.StatusOK, gin.H{"token": tokenStr, "expires_at": token.ExpiresAt()}, "Sign in successfully! ")
}
