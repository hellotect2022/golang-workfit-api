package user_service

import (
	"fmt"
	"go_server/models"
	"go_server/repository"
)

func AuthenticateUser(userId string, password string) (*models.User, error) {
	user, err := repository.FindUserByCredentials(userId, password)

	if err != nil {
		return nil, fmt.Errorf("사용자 인증 실패: %w", err)
	}

	fmt.Printf("사용자 인증 시도: ", user)

	return user, nil
}
