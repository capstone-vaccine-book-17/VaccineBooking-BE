package adminController

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// TODO CREATE ROLES
func (u *adminController) CreateRoles(c echo.Context) error {
	var payloads adminDto.RoleDTO

	if err := c.Bind(&payloads); err != nil {
		return err
	}
	fmt.Println(middleware.ClaimData(c, "roleID"))
	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	res, err := u.adminServ.CreateRoles(payloads)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "create role success",
		Code:    http.StatusOK,
		Data:    res,
	})
}
