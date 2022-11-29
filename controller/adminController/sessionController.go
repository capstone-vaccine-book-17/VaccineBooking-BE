package adminController

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// TODO CREATE SESSION
func (u *adminController) CreateSession(c echo.Context) error {
	medicalID, _ := middleware.ClaimData(c, "medicalID")

	conv_medicalID := medicalID.(float64)

	conv := uint(conv_medicalID)

	var payloads adminDto.SessionRequest

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	temp := adminDto.SessionRequest{
		Name:               payloads.Name,
		MedicalFacilitysId: conv,
		VaccineId:          payloads.VaccineId,
		StartTime:          payloads.StartTime,
		Kuota:              payloads.Kuota,
		Dosis:              payloads.Dosis,
		EndTime:            payloads.EndTime,
		Date:               payloads.Date,
	}

	res, err := u.adminServ.CreateSession(temp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "create session success",
		Code:    http.StatusOK,
		Data:    res,
	})
}

// TODO GET ALL SESSION
func (u *adminController) GetAllSession(c echo.Context) error {
	res, err := u.adminServ.GetAllSession()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "success",
		Code:    http.StatusOK,
		Data:    res,
	})
}

// TODO GET SESSION BY ID
func (u *adminController) GetSessionById(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	session := adminDto.SessionWithStatusDTO{
		SessionId: uint(convId),
	}

	res, err := u.adminServ.GetSessionById(session)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "success",
		Code:    http.StatusOK,
		Data:    res,
	})
}

// TODO UPDATE SESSION
func (u *adminController) UpdateSession(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	var payloads adminDto.SessionRequestUpdate

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	if err := c.Validate(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	session := adminDto.SessionRequestUpdate{
		SessionId: uint(convId),
		Name:      payloads.Name,
		VaccineId: payloads.VaccineId,
		Kuota:     payloads.Kuota,
		Dosis:     payloads.Dosis,
		StartTime: payloads.StartTime,
		EndTime:   payloads.EndTime,
	}

	res, err := u.adminServ.UpdateSession(session)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "success",
		Code:    http.StatusOK,
		Data:    res,
	})
}

// TODO DELETE SESSION
func (u *adminController) DeleteSession(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	payload := adminDto.SessionWithStatusDTO{
		SessionId: uint(convId),
	}

	err = u.adminServ.DeleteSession(payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "success",
		Code:    http.StatusOK,
	})
}
