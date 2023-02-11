package users

import (
	"errors"
	"net/http"
	"projek_fisioterapi/helper"
	"projek_fisioterapi/models"
	"projek_fisioterapi/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewUserServices(userRepo repositories.IUserRepository) IUserServices {
	return &userServices{IUserRepository: userRepo}
}

type IUserServices interface {
	GetProfile(id int) (models.User, error)
	Register(user models.User) error
}

type userServices struct {
	repositories.IUserRepository
}

func (s *userServices) Register(request models.User) error {
	var user models.User

	//check if email exist
	_, errCheckEmail := s.IUserRepository.GetUserByEmail(request.Email)
	if errors.Is(errCheckEmail, gorm.ErrRecordNotFound) {
		user.Name = request.Name
		user.Email = request.Email
		//hash password
		hashedPassword, errHashPassword := helper.HashPassword(request.Password)
		if errHashPassword != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, errHashPassword.Error())
		}
		user.Password = hashedPassword
		user.Phone = request.Phone
		user.Address = request.Address
		user.Role = request.Role
		user.Status = true
		user.IsAdmin = false

		//save user
		errSaveNewUser := s.IUserRepository.SaveNewUser(user)
		if errSaveNewUser != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, errSaveNewUser.Error())
		}
	} else if errCheckEmail != nil { //error other than not found
		return echo.NewHTTPError(http.StatusInternalServerError, errCheckEmail.Error())
	} else { //email has been taken
		return echo.NewHTTPError(http.StatusConflict, "Email has been taken")
	}

	return nil
}

func (s *userServices) GetProfile(id int) (models.User, error) {
	user, errGetProfile := s.IUserRepository.GetUserById(int(id))
	if errors.Is(errGetProfile, gorm.ErrRecordNotFound) {
		return models.User{}, echo.NewHTTPError(http.StatusNotFound, "Invalid JWT Data")
	} else if errGetProfile != nil {
		return models.User{}, echo.NewHTTPError(http.StatusInternalServerError, errGetProfile.Error())
	}

	result := user

	return result, nil
}
