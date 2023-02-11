package users

import (
	"net/http"
	"projek_fisioterapi/models"
	"projek_fisioterapi/services/users"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	users.IUserServices
}

func (h *UserHandler) Register(c echo.Context) error {
	var user models.User

	//bind
	errBind := c.Bind(&user)
	if errBind != nil {
		return errBind
	}

	//validasi
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return errValidate
	}

	//register
	err := h.IUserServices.Register(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User created",
	})
}

func (h *UserHandler) GetProfile(c echo.Context) error {
	// var user models.User
	// errBind := c.Bind(&user)
	// if errBind != nil {
	// 	return errBind
	// }

	// token, errDecodeJWT := helper.DecodeJWT(c)
	// if errDecodeJWT != nil {
	// 	return errDecodeJWT
	// }

	id := 1

	result, err := h.IUserServices.GetProfile(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
		"data":    result,
	})
}
