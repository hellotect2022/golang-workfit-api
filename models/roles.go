package models

type Roles struct {
	ID          uint   `gorm:"primaryKey"` // 역할 ID
	Name        string `gorm:"not null"`   // 역할 이름 // 관리자, 개발자, 보고자
	Description string `gorm:"type:text"`  // 역할 설명
}
