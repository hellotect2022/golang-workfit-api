package middleware

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func PostRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 요청 전 처리 (예: 로깅, 인증 등)
		log.Printf("요청 경로: %s, 메소드: %s", c.Request.URL.Path, c.Request.Method)

		// 요청 후 처리 (예: 응답 로깅 등)
		c.Next() // 다음 핸들러 실행
		log.Printf("응답 상태: %d", c.Writer.Status())

		// 에러 처리 (예: 에러 로깅 등)
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				log.Printf("에러 발생: %v", err)
			}
		}

		// 응답 후 처리
		fmt.Printf("요청 처리 완료: %s\n", c.Request.URL.Path)
	}
}
