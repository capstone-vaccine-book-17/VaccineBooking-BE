package adminController

import (
	"capstone_vaccine/utils"
	"net/http"

	"github.com/labstack/echo"
)

// TODO Get Profile
func (u *adminController) GetProfile(c echo.Context) error {

	res, err := u.adminServ.GetProfile()
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Profile",
		Code:    http.StatusOK,
		Data:    res,
	})
}
