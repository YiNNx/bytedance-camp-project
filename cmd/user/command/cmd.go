/*
* @Author: Go不浪队
* @Date: 2023-02-07 01:26:16
 * @LastEditTime: 2023-02-07 01:50:03
* @Description:
*/

package command

import "tiktok/db"

func CreateUser(c *db.Context,u *db.User)(id uint,err error ){
	res:=c.Transaction.Create(u)
	return u.ID ,res.Error
}

func UsernameDuplicate(c *db.Context,username string) (bool, error) {
	// TODO:
	return false, nil
}

func FindUser(c *db.Context,query *db.User) (*db.User, error) {
	res := &db.User{}
	err := c.Transaction.Where(query).Take(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

