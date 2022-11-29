package adminController

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/utils"
	"net/http"

	"github.com/labstack/echo"
)

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
