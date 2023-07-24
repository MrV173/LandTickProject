package ticketdto

import "landtick/models"

type CreateTicketRequest struct {
	Name_Train    string `json:"name_train" form:"name_train" validate:"required"`
	Type_Train    string `json:"type_train" form:"type_train" validate:"required"`
	Start_Date    string `json:"start_date" form:"start_date" validate:"required"`
	StationID     int    `json:"station_id" form:"station_id"`
	Start_Time    string `json:"start_time" form:"start_time" validate:"required"`
	DestinationID int    `json:"destination_id" form:"destination_id"`
	Arrival_Time  string `json:"arrival_time" form:"arrival_time" validate:"required"`
	Price         int    `json:"price" form:"price" validate:"required"`
	Qty           string `json:"qty" form:"qty" validate:"required"`
	UserID        int    `json:"user_id"`
}

type UpdateTicketRequest struct {
	Name_Train    string `json:"name_train" form:"name_train"`
	Type_Train    string `json:"type_train" form:"type_train"`
	Start_Date    string `json:"start_date" form:"start_date"`
	StationID     int    `json:"station_id" form:"station_id"`
	Start_Time    string `json:"start_time" form:"start_time"`
	DestinationID int    `json:"destination_id"`
	Arrival_Time  string `json:"arrival_time" form:"arrival_time"`
	Price         int    `json:"price" form:"price"`
	Qty           string `json:"qty" form:"qty"`
	UserID        int    `json:"user_id"`
}

type TicketResponse struct {
	ID            int                        `json:"id"`
	Name_Train    string                     `json:"name_train"`
	Type_Train    string                     `json:"type_train"`
	Start_Date    string                     `json:"start_date"`
	StationID     int                        `json:"station_id"`
	Station       models.StationResponse     `json:"station"`
	Start_Time    string                     `json:"start_time" `
	DestinationID int                        `json:"destination_id"`
	Destination   models.DestinationResponse `json:"destination"`
	Arrival_Time  string                     `json:"arrival_time"`
	Price         int                        `json:"price"`
	Qty           string                     `json:"qty"`
	UserID        int                        `json:"user_id"`
	User          models.UserResponse        `json:"user"`
}
