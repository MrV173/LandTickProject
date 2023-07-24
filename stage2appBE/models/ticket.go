package models

type Ticket struct {
	ID            int                 `json:"id" gorm:"primary_key:auto_increment"`
	Name_Train    string              `json:"name_train" gorm:"type: varchar(255)"`
	Type_Train    string              `json:"type_train" gorm:"type: varchar(255)"`
	Start_Date    string              `json:"start_date" gorm:"type: varchar(255)"`
	StationID     int                 `json:"station_id"`
	Station       StationResponse     `json:"station"`
	Start_Time    string              `json:"start_time" gorm:"type:string"`
	DestinationID int                 `json:"destination_id"`
	Destination   DestinationResponse `json:"destination"`
	Arrival_Time  string              `json:"arrival_time" gorm:"type:string"`
	Price         int                 `json:"price" gorm:"type: int"`
	Qty           string              `json:"qty" gorm:"type:int"`
	UserID        int                 `json:"user_id"`
	User          UserResponse        `json:"user"`
}

type TicketResponse struct {
	ID            int    `json:"id"`
	Name_Train    string `json:"name_train"`
	Type_Train    string `json:"type_train"`
	Start_Date    string `json:"start_date"`
	StationID     int    `json:"station_id"`
	Start_Time    string `json:"start_time"`
	DestinationID int    `json:"destination_id"`
	Arrival_Time  string `json:"arrival_time"`
	Price         int    `json:"price"`
	Qty           string `json:"qty"`
	UserID        int    `json:"user_id"`
}

func (TicketResponse) TableName() string {
	return "tickets"
}
