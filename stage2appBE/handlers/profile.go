package handlers

import (
	profiledto "landtick/dto/profile"
	dto "landtick/dto/result"
	"landtick/models"
	"landtick/repository"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerProfile struct {
	ProfileRepository repository.ProfielRepository
}

func ProfileHandler(profileRepository repository.ProfielRepository) *handlerProfile {
	return &handlerProfile{profileRepository}
}

func (h *handlerProfile) FindProfile(c echo.Context) error {
	profiles, err := h.ProfileRepository.FindProfile()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: profiles,
	})

}

func (h *handlerProfile) GetProfile(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	profile, err := h.ProfileRepository.GetProfile(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertProfileResponse(profile),
	})

}

func (h *handlerProfile) CreateProfile(c echo.Context) error {
	request := new(profiledto.CreateProfileRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest, Message: err.Error()})
	}

	profile := models.Profile{
		FullName:      request.FullName,
		UserName:      request.UserName,
		Email:         request.Email,
		Password:      request.Password,
		No_Telphone:   request.No_Telphone,
		Jenis_Kelamin: request.Jenis_Kelamin,
		Alamat:        request.Alamat,
		UserID:        request.UserID,
	}

	response, err := h.ProfileRepository.CreateProfile(profile)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertProfileResponse(response),
	})
}

func (h *handlerProfile) DeleteProfile(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	profile, err := h.ProfileRepository.GetProfile(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	data, err := h.ProfileRepository.DeleteProfile(profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertProfileResponse(data),
	})
}

func (h *handlerProfile) UpdateProfile(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	profile, err := h.ProfileRepository.GetProfile(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	request := new(profiledto.UpdateProfileRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if request.FullName != "" {
		profile.FullName = request.FullName
	}

	if request.UserName != "" {
		profile.UserName = request.UserName
	}

	if request.Email != "" {
		profile.Email = request.Email
	}

	if request.Password != "" {
		profile.Password = request.Password
	}

	if request.Jenis_Kelamin != "" {
		profile.Jenis_Kelamin = request.Jenis_Kelamin
	}

	if request.Alamat != "" {
		profile.Alamat = request.Alamat
	}

	response, err := h.ProfileRepository.UpdateProfile(profile)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: response,
	})

}

func convertProfileResponse(u models.Profile) profiledto.ProfileResponse {
	return profiledto.ProfileResponse{
		FullName:      u.FullName,
		UserName:      u.UserName,
		Email:         u.Email,
		Password:      u.Password,
		No_Telphone:   u.No_Telphone,
		Jenis_Kelamin: u.Jenis_Kelamin,
		Alamat:        u.Alamat,
		UserID:        u.UserID,
		User:          u.User,
	}
}
