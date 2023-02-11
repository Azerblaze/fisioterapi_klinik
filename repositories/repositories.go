package repositories

import (
	"projek_fisioterapi/models"
)

type IUserRepository interface {
	SaveNewUser(user models.User) error
	GetUserById(userId int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
}
