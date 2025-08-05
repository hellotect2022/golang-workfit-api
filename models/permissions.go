package models

import "fmt"

type Permissions struct {
	ID          uint   `gorm:"primaryKey"` // 권한 ID
	Name        string `gorm:"not null"`   // 권한 이름 // 읽기, 쓰기, 수정, 삭제 issue_view, issue_create, issue_edit, issue_delete
	Description string `gorm:"type:text"`  // 권한 설명
}

func (p *Permissions) TableName() string {
	return "permissions"
}
func (p *Permissions) GetID() uint {
	return p.ID
}
func (p *Permissions) GetName() string {
	return p.Name
}
func (p *Permissions) GetDescription() string {
	return p.Description
}
func (p *Permissions) SetID(id uint) {
	p.ID = id
}
func (p *Permissions) SetName(name string) {
	p.Name = name
}
func (p *Permissions) SetDescription(description string) {
	p.Description = description
}
func (p *Permissions) IsValid() bool {
	return p.ID > 0 && p.Name != ""
}
func (p *Permissions) String() string {
	return p.Name + ": " + p.Description
}
func (p *Permissions) Equals(other *Permissions) bool {
	if other == nil {
		return false
	}
	return p.ID == other.ID && p.Name == other.Name && p.Description == other.Description
}
func (p *Permissions) Clone() *Permissions {
	return &Permissions{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
	}
}
func (p *Permissions) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("permission name cannot be empty")
	}
	if p.Description == "" {
		return fmt.Errorf("permission description cannot be empty")
	}
	return nil
}
func (p *Permissions) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          p.ID,
		"name":        p.Name,
		"description": p.Description,
	}
}
func (p *Permissions) FromMap(data map[string]interface{}) error {
	if id, ok := data["id"].(uint); ok {
		p.ID = id
	} else {
		return fmt.Errorf("invalid or missing ID")
	}
	if name, ok := data["name"].(string); ok {
		p.Name = name
	} else {
		return fmt.Errorf("invalid or missing Name")
	}
	if description, ok := data["description"].(string); ok {
		p.Description = description
	} else {
		return fmt.Errorf("invalid or missing Description")
	}
	return nil
}
