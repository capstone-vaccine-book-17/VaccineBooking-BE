package citizenController

import (
	"capstone_vaccine/middleware"
	"capstone_vaccine/utils"
	"net/http"

	"github.com/labstack/echo"
)

func (u *citizenController) GetTicket(c echo.Context) error {
	citizenID, _ := middleware.ClaimData(c, "citizenID")
	conv_citizenID := citizenID.(float64)
	conv := uint(conv_citizenID)

	res, err := u.citizenServ.GetTicket(conv)

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
