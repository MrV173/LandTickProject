package handlers

import (
	dto "landtick/dto/result"
	ticketdto "landtick/dto/ticket"
	"landtick/models"
	"landtick/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type handlerTicket struct {
	TicketRepository repository.TicketRepository
}

func TicketHandler(ticketRepository repository.TicketRepository) *handlerTicket {
	return &handlerTicket{ticketRepository}
}

func (h *handlerTicket) FindTicket(c echo.Context) error {
	tickets, err := h.TicketRepository.FindTicket()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: tickets,
	})
}

func (h *handlerTicket) GetTicket(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ticket, err := h.TicketRepository.GetTicket(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertTicketResponse(ticket)})
}

func (h *handlerTicket) CreateTicket(c echo.Context) error {
	request := new(ticketdto.CreateTicketRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	ticket := models.Ticket{
		Name_Train:    request.Name_Train,
		Type_Train:    request.Type_Train,
		Start_Date:    request.Start_Date,
		StationID:     request.StationID,
		Start_Time:    request.Start_Time,
		DestinationID: request.DestinationID,
		Arrival_Time:  request.Arrival_Time,
		Price:         request.Price,
		Qty:           request.Qty,
		UserID:        int(userId),
	}

	response, err := h.TicketRepository.CreateTicket(ticket)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertTicketResponse(response)})

}

func (h *handlerTicket) DeleteTicket(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ticket, err := h.TicketRepository.GetTicket(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TicketRepository.DeleteTicket(ticket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertTicketResponse(data),
	})
}

func (h *handlerTicket) UpdateTicket(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ticket, err := h.TicketRepository.GetTicket(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	request := new(ticketdto.UpdateTicketRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	if request.Name_Train != "" {
		ticket.Name_Train = request.Name_Train
	}

	if request.Type_Train != "" {
		ticket.Type_Train = request.Type_Train
	}

	if request.Start_Date != "" {
		ticket.Start_Date = request.Start_Date
	}

	if request.Start_Time != "" {
		ticket.Start_Time = request.Start_Time
	}

	if request.Arrival_Time != "" {
		ticket.Arrival_Time = request.Arrival_Time
	}

	if request.Qty != "" {
		ticket.Qty = request.Qty
	}
	response, err := h.TicketRepository.UpdateTicket(ticket)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response})

}

func convertTicketResponse(u models.Ticket) ticketdto.TicketResponse {
	return ticketdto.TicketResponse{
		ID:            u.ID,
		Name_Train:    u.Name_Train,
		Type_Train:    u.Type_Train,
		Start_Date:    u.Start_Date,
		StationID:     u.StationID,
		Station:       u.Station,
		Start_Time:    u.Start_Time,
		DestinationID: u.DestinationID,
		Destination:   u.Destination,
		Arrival_Time:  u.Arrival_Time,
		Price:         u.Price,
		Qty:           u.Qty,
		UserID:        u.UserID,
		User:          u.User,
	}
}
