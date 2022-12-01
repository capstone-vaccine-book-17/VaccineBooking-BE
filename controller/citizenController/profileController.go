package citizenController

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/utils"
	"net/http"

	"github.com/labstack/echo"
)

// TODO Get Profile
func (u *citizenController) GetProfile(c echo.Context) error {

	citizenID, _ := middleware.ClaimData(c, "citizenID")

	conv_citizenID := citizenID.(float64)

	conv := uint(conv_citizenID)

	temp := citizenDto.ProfileReq{
		CitizenID: conv,
	}

	res, err := u.citizenServ.GetProfile(temp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "MY PROFILE",
		Code:    http.StatusOK,
		Data:    res,
	})

}
