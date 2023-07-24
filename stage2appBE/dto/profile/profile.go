package profiledto

import "landtick/models"

type CreateProfileRequest struct {
	FullName      string `json:"fullname"`
	UserName      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	No_Telphone   int    `json:"no_telphone"`
	Jenis_Kelamin string `json:"jenis_kelamin"`
	Alamat        string `json:"alamat"`
	UserID        int    `json:"user_id"`
}

type UpdateProfileRequest struct {
	FullName      string `json:"fullname" form:"fullname"`
	UserName      string `json:"username" form:"username"`
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	No_Telphone   int    `json:"no_telphone" form:"no_telphone"`
	Jenis_Kelamin string `json:"jenis_kelamin" form:"jenis_kelamin"`
	Alamat        string `json:"alamat" form:"alamat"`
	UserID        int    `json:"user_id" form:"user_id"`
}

type ProfileResponse struct {
	FullName      string              `json:"fullname"`
	UserName      string              `json:"username"`
	Email         string              `json:"email"`
	Password      string              `json:"password"`
	No_Telphone   int                 `json:"no_telphone"`
	Jenis_Kelamin string              `json:"jenis_kelamin"`
	Alamat        string              `json:"alamat"`
	UserID        int                 `json:"user_id"`
	User          models.UserResponse `json:"user"`
}
