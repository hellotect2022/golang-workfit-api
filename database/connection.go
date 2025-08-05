package database

import (
	"fmt"
	"go_server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB // 전역 DB 변수 포인터로 선언

func Connect() (*gorm.DB, error) {
	dsn := config.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 테이블 이름을 단수형으로 (User → user)
			NameReplacer:  nil,  // 선택: 단어 치환기
			TablePrefix:   "",   // 선택: 테이블 접두사
		},
	})
	if err != nil {
		return nil, fmt.Errorf("DB 연결 실패: %w", err)
	}

	DB = db // 전역 변수에 DB 할당

	return db, nil
}
