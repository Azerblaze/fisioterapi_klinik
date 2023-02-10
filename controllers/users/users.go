package users

import (
	"net/http"
	"projek_fisioterapi/services/users"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	users.IUserServices
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
