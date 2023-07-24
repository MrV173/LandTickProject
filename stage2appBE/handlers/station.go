package handlers

import (
	dto "landtick/dto/result"
	stationdto "landtick/dto/station"
	"landtick/models"
	"landtick/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerStation struct {
	StationRepository repository.StationRepository
}

func StationHandler(stationRepository repository.StationRepository) *handlerStation {
	return &handlerStation{stationRepository}
}

func (h *handlerStation) FindStation(c echo.Context) error {
	stations, err := h.StationRepository.FindStation()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: stations,
	})
}

func (h *handlerStation) GetStation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	station, err := h.StationRepository.GetStation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: StationConvertResponse(station)})
}

func (h *handlerStation) CreateStation(c echo.Context) error {
	request := new(stationdto.CreateStationRequest)
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

	station := models.Station{
		Name: request.Name,
	}

	response, err := h.StationRepository.CreateStation(station)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: StationConvertResponse(response)})

}

func (h *handlerStation) DeleteStation(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	station, err := h.StationRepository.GetStation(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: StationConvertResponse(station),
	})
}

func StationConvertResponse(u models.Station) stationdto.StationResponse {
	return stationdto.StationResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}
