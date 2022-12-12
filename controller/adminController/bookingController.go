package adminController

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// TODO CREATE BOOKING
func (u *adminController) CreateBooking(c echo.Context) error {
	var payloads adminDto.BookingDto

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	res, err := u.adminServ.CreateBooking(payloads)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "create booking success",
		Code:    http.StatusOK,
		Data:    res,
	})
}

// TODO UPDATE BOOKING
func (u *adminController) UpdateBooking(c echo.Context) error {

	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	var payloads adminDto.UpdateBooking

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	temp := adminDto.UpdateBooking{
		BookingId: uint(convId),
		Status:    payloads.Status,
	}

	res, err := u.adminServ.UpdateBooking(temp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "update booking success",
		Code:    http.StatusOK,
		Data:    res,
	})
}

// TODO GET ALL BOOKING
func (u *adminController) GetAllBooking(c echo.Context) error {
	medicalID, _ := middleware.ClaimData(c, "medicalID")

	conv_medicalID := medicalID.(float64)

	conv := uint(conv_medicalID)
	res, err := u.adminServ.GetAllBooking(conv)

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

// TODO GET BOOKING BY ID
func (u *adminController) GetBookingById(c echo.Context) error {

	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	payloads := adminDto.BookingAllDto{
		BookingId: uint(convId),
	}

	res, err := u.adminServ.GetBookingById(payloads)

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

// TODO DELETE BOOKING
func (u *adminController) DeleteBooking(c echo.Context) error {

	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	payloads := adminDto.BookingAllDto{
		BookingId: uint(convId),
	}

	errS := u.adminServ.DeleteBooking(payloads)

	if errS != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: errS.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "success",
		Code:    http.StatusOK,
	})
}
