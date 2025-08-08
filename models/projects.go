package models

import "time"

type Projects struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`       // 프로젝트 이름
	Description string    `gorm:"type:text"`      // 프로젝트 설명
	IsPublic    bool      `gorm:"default:false"`  // 공개 여부
	CreatedAt   time.Time `gorm:"autoCreateTime"` // 생성 시 자동 설정
	UpdatedAt   time.Time `gorm:"autoUpdateTime"` // 수정 시 자동 갱신
	OwnerID     string    `gorm:"not null"`       // 소유자 ID (User 모델과의 관계를 나타냄)
	ParentID    uint      // 부모 프로젝트 ID
}
