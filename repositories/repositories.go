package repositories

import (
	"projek_fisioterapi/models"
)

type IUserRepository interface {
	GetUserById(userId int) (models.User, error)
}
