package models

type Destination struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" gorm:"type: varchar(255)"`
}

type DestinationResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (DestinationResponse) TableName() string {
	return "destinations"
}
