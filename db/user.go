/*
 * @Author: Go不浪队
 * @Date: 2023-02-06 22:32:06
 * @LastEditTime: 2023-02-06 22:38:06
 * @Description:
 */

package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	PasswordHash       string

	Username     string
	FollowCount int64
	FollowerCount int64
}
