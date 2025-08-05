package database

import (
	"fmt"
	"go_server/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// 마이그레이션을 수행합니다.
	// 테이블 관계 요약
	// User <-> ProjectMemberships             1:N
	// Projects <-> ProjectMemberships         1:N
	// Roles <-> ProjectMemberships            1:N
	// Roles <-> RolePermissions               1:N
	// Permissions <-> RolePermissions         1:N

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("마이그레이션 실패: %w", err)
	}

	// 마이그레이션을 수행합니다.
	if err := db.AutoMigrate(&models.Roles{}); err != nil {
		return fmt.Errorf("마이그레이션 실패: %w", err)
	}

	// 마이그레이션을 수행합니다.
	if err := db.AutoMigrate(&models.RolePermissions{}); err != nil {
		return fmt.Errorf("마이그레이션 실패: %w", err)
	}

	// 마이그레이션을 수행합니다.
	if err := db.AutoMigrate(&models.Projects{}); err != nil {
		return fmt.Errorf("마이그레이션 실패: %w", err)
	}

	// 마이그레이션을 수행합니다.
	if err := db.AutoMigrate(&models.ProjectMemberships{}); err != nil {
		return fmt.Errorf("마이그레이션 실패: %w", err)
	}

	// 마이그레이션을 수행합니다.
	if err := db.AutoMigrate(&models.Permissions{}); err != nil {
		return fmt.Errorf("마이그레이션 실패: %w", err)
	}

	return nil
}
