package models

/*
사용자, 프로젝트, 그리고 권한 을 연결하는 핵식 테이블
*/
type ProjectMemberships struct {
	ID        uint `gorm:"primaryKey"` // 프로젝트 멤버십 ID
	UserID    uint `gorm:"not null"`   // 사용자 ID (User 모델과의 관계를 나타냄)
	ProjectID uint `gorm:"not null"`   // 프로젝트 ID (Projects 모델과의 관계를 나타냄)
	RoleID    uint `gorm:"not null"`   // 역할 ID (Roles 모델과의 관계를 나타냄)

	// 관계 설정
	User     User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Roles    Roles    `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Projects Projects `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
