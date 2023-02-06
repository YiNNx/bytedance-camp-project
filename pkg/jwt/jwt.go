/*
 * @Author: Go不浪队
 * @Date: 2023-02-05 20:35:14
 * @LastEditTime: 2023-02-06 21:36:18
 * @Description:
 */

package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"

	"tiktok/pkg/config"
)

type JwtUserClaims struct {
	Id string `json:"id"` // weixin open id
	jwt.StandardClaims
}

func GenerateToken(id string) (string, error) {
	claims := &JwtUserClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 7 * 24).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(config.C.JWT.Secret))
}

