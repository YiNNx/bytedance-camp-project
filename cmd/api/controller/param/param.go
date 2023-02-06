/*
 * @Author: Go不浪队
 * @Date: 2023-02-05 20:23:08
 * @LastEditTime: 2023-02-07 00:20:16
 * @Description:
 */

package param

type ReqSignUp struct {
	Username string `form:"username" binding:"required,max=32"`
	Password string `form:"password" binding:"required,max=32"`
}

type ResSignUp struct {
	UserId int64    `json:"user_id"`
	Token  string `json:"token"`
}

type ReqSignIn struct {
	Username string `form:"username" binding:"required,max=32"`
	Password string `form:"password" binding:"required,max=32"`
}
 