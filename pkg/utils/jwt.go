package util

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("Python")

type Claims struct {
	ID        uint   `json:"id"`
	Username  string `json"username"`
	Authority int    `json:"authority"`
	Isrefresh bool   `json:"isrefresh"`
	jwt.StandardClaims
}

func GenerateToken(claims Claims) (string, error) {
	// nowTime := time.Now()
	// expireTime := nowTime.Add(24 *time.Hour)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) { return jwtSecret, nil })
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func MyGenerateToken(user_id uint, username string, authority int, need_refresh_token bool) (string, string, error) {

	// 生成业务token 2h有效期
	expiry := time.Now().Add(2 * time.Hour)

	claims := Claims{
		ID:        user_id,
		Username:  username,
		Authority: authority, // 权限用户
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry.Unix(),
			Issuer:    "go_ctry",
		},
	}

	// 调用生成 token
	token, err := GenerateToken(claims)

	var refreshToken string
	if need_refresh_token {
		// 生成refresh_token 14天有效期
		reEpiry := time.Now().Add(14 * time.Hour)

		refreshClaims := Claims{
			ID:        user_id,
			Username:  username,
			Authority: authority, // 权限用户
			Isrefresh: true,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: reEpiry.Unix(),
				Issuer:    "go_ctry",
			},
		}
		refreshToken, err = GenerateToken(refreshClaims)
	}
	return token, refreshToken, err

}
