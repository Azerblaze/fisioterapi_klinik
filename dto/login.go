package dto

type Login struct {
	Email    string `json:"username" form:"username" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}
