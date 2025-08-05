package handler

import (
	"go_server/service/token_service"
	"go_server/service/user_service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginRequest struct {
		UserId   string `json:"userId"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// 사용자 검증 서비스 호출
	result, err := user_service.AuthenticateUser(loginRequest.UserId, loginRequest.Password)

	if err != nil {
		c.JSON(401, gin.H{"error": "Authentication failed", "details": err.Error()})
		return
	}

	// 토큰 생성 서비스 호출
	token, err := token_service.GenerateToken(result.UserId)

	if err != nil {
		c.JSON(500, gin.H{"error": "Token generation failed", "details": err.Error()})
		return
	}

	// 예: 데이터베이스에서 사용자 인증, JWT 토큰 생성 등
	// 현재는 성공 응답을 반환합니다.
	c.JSON(200, gin.H{"message": "Login successful", "token": token})
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Test successful"})
}
