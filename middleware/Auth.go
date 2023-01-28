package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

var (
	Secret = []byte("Gegege")
	// TokenExpireDuration = time.Hour * 2 过期时间
)

type JWTClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

// 生成token
func GenToken(userid int64, userName string) (string, error) {
	claims := JWTClaims{
		UserId:   userid,
		Username: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "douyin",
			//ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),可用于设定token过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("Gegege"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// 解析token
func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Invalid token")
}

// 验证token
func VerifyToken(tokenString string) (int64, error) {
	if tokenString == "" {
		return int64(0), nil
	}
	claims, err := ParseToken(tokenString)
	if err != nil {
		return int64(0), err
	}
	return claims.UserId, nil
}
