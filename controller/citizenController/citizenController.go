package citizenController

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/service/citizenService"
	"capstone_vaccine/utils"
	"net/http"

	"github.com/labstack/echo"
)

type CitizenController interface{}

type citizenController struct {
	citizenServ citizenService.CitizenService
}

func NewCitizenController(citizenService citizenService.CitizenService) *citizenController {
	return &citizenController{
		citizenServ: citizenService,
	}
}

// TODO LOGIN
func (u *citizenController) LoginCitizen(c echo.Context) error {
	var payloads citizenDto.LoginDto

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	res, err := u.citizenServ.LoginCitizen(payloads)

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

// TODO REGISTER
func (u *citizenController) RegisterCitizen(c echo.Context) error {
	var payloads citizenDto.RegisterDto

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	res, err := u.citizenServ.RegisterCitizen(payloads)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "register success",
		Code:    http.StatusOK,
		Data:    res,
	})
}
