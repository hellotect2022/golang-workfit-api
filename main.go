package main

import (
	"go_server/database"
	"go_server/handler"
	"go_server/middleware"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// "gorm.io/driver/postgres"
// "gorm.io/gorm"

func main() {
	db, err := database.Connect()

	if err != nil {
		log.Fatalf("DB 연결 실패: %v", err)
		return
	}

	database.Migrate(db)
	log.Println("DB 연결 및 마이그레이션 성공")

	// gin 서버 시작

	r := gin.Default()
	r.Use(middleware.PostRequest())
	// CORS 설정
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/api/login", handler.Login)

	r.Use(middleware.JWTAuth())

	r.GET("/api/test", handler.Test)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("서버 시작 실패: %v", err)
		return
	}
}
