package repository

import (
	"go_server/database"
	"go_server/models"
	"go_server/utils"

	"gorm.io/gorm"
)

func FindUserByCredentials(userId, password string) (*models.User, error) {
	var user models.User

	result := database.DB.Where("user_id = ?", userId).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if ok, err := utils.CheckPasswordHash(password, user.PasswordHash); err != nil || !ok {
		return nil, gorm.ErrRecordNotFound // 비밀번호가 일치하지 않으면
	}

	return &user, nil
}

func CreateUser(db *gorm.DB, user *models.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
func GetUserByID(db *gorm.DB, id uint) (*models.User, error) {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func UpdateUser(db *gorm.DB, user *models.User) error {
	if err := db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
func DeleteUser(db *gorm.DB, id uint) error {
	if err := db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
