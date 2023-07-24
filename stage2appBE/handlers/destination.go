package handlers

import (
	destinationdto "landtick/dto/destination"
	dto "landtick/dto/result"
	"landtick/models"
	"landtick/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerDestination struct {
	DestinationRepository repository.DestinationRepository
}

func DestinationHandler(destinationRepository repository.DestinationRepository) *handlerDestination {
	return &handlerDestination{destinationRepository}
}

func (h *handlerDestination) FindDestination(c echo.Context) error {
	destinations, err := h.DestinationRepository.FindDestination()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: destinations,
	})
}

func (h *handlerDestination) GetDestination(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	destination, err := h.DestinationRepository.GetDestination(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: DestinationConvertResponse(destination)})
}

func (h *handlerDestination) CreateDestination(c echo.Context) error {
	request := new(destinationdto.CreateDestinationRequest)
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

	destination := models.Destination{
		Name: request.Name,
	}

	response, err := h.DestinationRepository.CreateDestination(destination)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: DestinationConvertResponse(response)})

}

func (h *handlerDestination) DeleteDestination(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	destination, err := h.DestinationRepository.GetDestination(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: DestinationConvertResponse(destination),
	})
}

func DestinationConvertResponse(u models.Destination) destinationdto.DestinationResponse {
	return destinationdto.DestinationResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}
