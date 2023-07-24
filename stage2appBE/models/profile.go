package models

type Profile struct {
	ID            int          `json:"id" gorm:"primaryKey:autoIncrement"`
	FullName      string       `json:"fullname" gorm:"type: varchar(255)"`
	UserName      string       `json:"username" gorm:"type: varchar(255)"`
	Email         string       `json:"email" gorm:"type: varchar(255)"`
	Password      string       `json:"password" gorm:"type: varchar(255)"`
	No_Telphone   int          `json:"no_telphone" gorm:"type: int"`
	Jenis_Kelamin string       `json:"jenis_kelamin" gorm:"type: varchar(255)"`
	Alamat        string       `json:"alamat" gorm:"type: varchar(255)"`
	UserID        int          `json:"user_id"`
	User          UserResponse `json:"user"`
}

type ProfileResponse struct {
	FullName      string `json:"fullname"`
	UserName      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	No_Telphone   int    `json:"no_telphone"`
	Jenis_Kelamin string `json:"jenis_kelamin"`
	Alamat        string `json:"alamat"`
	UserID        int    `json:"user_id"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
