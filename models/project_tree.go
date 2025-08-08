package models

type ProjectTree struct {
	ID          uint   `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	ParentID    *uint  `gorm:"column:parent_id" json:"parentId"` // NULL이 가능하므로 포인터 사용
	IsPublic    bool   `gorm:"column:is_public" json:"isPublic"`
	Description string `gorm:"column:description" json:"description"`
	OwnerID     string `gorm:"column:owner_id" json:"ownerId"`
	Level       int    `gorm:"column:level" json:"level"`
}
