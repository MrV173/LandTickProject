package models

type User struct {
	ID          int                 `json:"id" gorm:"primaryKey:autoIncrement"`
	FullName    string              `json:"fullname" gorm:"type: varchar(255)"`
	UserName    string              `json:"username" gorm:"type: varchar(255)"`
	Email       string              `json:"email" gorm:"type: varchar(255)"`
	Password    string              `json:"password" gorm:"type: varchar(255)"`
	Role        string              `json:"role" gorm:"type : varchar(255)"`
	Profile     ProfileResponse     `json:"profile"`
	Transaction TransactionResponse `json:"transaction"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

func (UserResponse) TableName() string {
	return "users"
}
