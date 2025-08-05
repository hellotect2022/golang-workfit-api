package models

import "fmt"

type RolePermissions struct {
	Role         string // roles 테이블과 permissions 테이블을 연결하는 역할
	RoleID       uint   `gorm:"not null"` // 역할 ID
	PermissionID uint   `gorm:"not null"` // 권한 ID

	// 관계 설정
	Roles       Roles       `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Permissions Permissions `gorm:"foreignKey:PermissionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (rp *RolePermissions) TableName() string {
	return "role_permissions"
}
func (rp *RolePermissions) GetRoleID() uint {
	return rp.RoleID
}
func (rp *RolePermissions) GetPermissionID() uint {
	return rp.PermissionID
}
func (rp *RolePermissions) SetRoleID(roleID uint) {
	rp.RoleID = roleID
}
func (rp *RolePermissions) SetPermissionID(permissionID uint) {
	rp.PermissionID = permissionID
}
func (rp *RolePermissions) IsValid() bool {
	return rp.RoleID > 0 && rp.PermissionID > 0
}
func (rp *RolePermissions) String() string {
	return "Role: " + rp.Role + ", RoleID: " + string(rp.RoleID) + ", PermissionID: " + string(rp.PermissionID)
}
func (rp *RolePermissions) Equals(other *RolePermissions) bool {
	if other == nil {
		return false
	}
	return rp.RoleID == other.RoleID && rp.PermissionID == other.PermissionID
}
func (rp *RolePermissions) Clone() *RolePermissions {
	return &RolePermissions{
		Role:         rp.Role,
		RoleID:       rp.RoleID,
		PermissionID: rp.PermissionID,
	}
}
func (rp *RolePermissions) Validate() error {
	if rp.RoleID == 0 {
		return fmt.Errorf("role ID cannot be zero")
	}
	if rp.PermissionID == 0 {
		return fmt.Errorf("permission ID cannot be zero")
	}
	return nil
}
func (rp *RolePermissions) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"role":          rp.Role,
		"role_id":       rp.RoleID,
		"permission_id": rp.PermissionID,
	}
}
func (rp *RolePermissions) FromMap(data map[string]interface{}) error {
	if role, ok := data["role"].(string); ok {
		rp.Role = role
	} else {
		return fmt.Errorf("role is required")
	}
	if roleID, ok := data["role_id"].(uint); ok {
		rp.RoleID = roleID
	} else {
		return fmt.Errorf("role ID is required")
	}
	if permissionID, ok := data["permission_id"].(uint); ok {
		rp.PermissionID = permissionID
	} else {
		return fmt.Errorf("permission ID is required")
	}
	return nil
}
