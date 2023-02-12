package users

import (
	"projek_fisioterapi/models"
	"projek_fisioterapi/repositories"

	"gorm.io/gorm"
)

type DBGorm struct {
	DB *gorm.DB
}

func NewGorm(db *gorm.DB) repositories.IUserRepository {
	return &DBGorm{
		DB: db,
	}
}

func (db DBGorm) SaveNewUser(user models.User) error {
	result := db.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db DBGorm) GetUserById(userId int) (models.User, error) {
	var user models.User
	err := db.DB.
		Where("id = ?", userId).
		First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (db DBGorm) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
