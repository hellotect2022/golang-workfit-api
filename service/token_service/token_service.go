package token_service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId string) (string, error) {
	// 여기에 JWT 토큰 생성 로직을 구현합니다.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   userId,
		"userRole": userId,
		"exp":      jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24시간 후 만료
	})

	fmt.Println("토큰 생성:", token)

	tokenString, err := token.SignedString([]byte("your_secret_key")) // 비밀 키는 안전하게 관리해야 합니다.

	if err != nil {
		return "", err
	}

	fmt.Println("토큰 문자열:", tokenString)

	return tokenString, nil

}
