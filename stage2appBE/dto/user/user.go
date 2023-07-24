package userdto

import "landtick/models"

type CreateUserRequest struct {
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	UserName string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateUserRequest struct {
	FullName string `json:"fullname" form:"fullname"`
	UserName string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	ID          int                        `json:"id"`
	FullName    string                     `json:"fullname"`
	UserName    string                     `json:"username"`
	Email       string                     `json:"email"`
	Profile     models.ProfileResponse     `json:"profile"`
	Transaction models.TransactionResponse `json:"Transaction"`
}
