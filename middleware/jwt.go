package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("JWT 인증 미들웨어 실행")
		c.Next() // JWT 인증 로직을 여기에 구현할 수 있습니다.
		fmt.Printf("요청 처리 완료: %s\n", c.Request.URL.Path)
		fmt.Println("JWT 인증 미들웨어 종료")
	}
}
