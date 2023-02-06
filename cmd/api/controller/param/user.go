/*
 * @Author: Go不浪队
 * @Date: 2023-02-05 20:23:08
 * @LastEditTime: 2023-02-06 20:23:08
 * @Description:
 */

package param

type ReqSignUp struct {
	Username string `json:"username" binding:"required,min=5,max=32"`
	Email    string `json:"email" binding:"required,max=128,email"`
	Password string `json:"password" binding:"required,min=6,max=64"`
	Intro    string `json:"intro" binding:"max=512"`
	Github   string `json:"github" binding:"max=39"`
	School   string `json:"school" binding:"max=32"`
	Website  string `json:"website" binding:"max=128,url"`
}
 
type ReqSignIn struct {
	Username string `json:"username" binding:"required,min=5,max=32"`
	Password string `json:"password" binding:"required,min=6,max=64"`
}
 