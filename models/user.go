package models

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	UserId       string `gorm:"uniqueIndex;not null"` // ✅ 고유 인덱스 추가
	Username     string
	Email        string
	PasswordHash string
	IsActive     bool
	CreatedAt    time.Time `gorm:"autoCreateTime"` // ✅ 생성 시 자동 설정
	UpdatedAt    time.Time `gorm:"autoUpdateTime"` // ✅ 수정 시 자동 갱신
}
