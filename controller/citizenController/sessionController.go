package citizenController

import (
	"capstone_vaccine/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// TODO GET SESSION BY MEDICAL FACILITYS ID
func (u *citizenController) GetSessionByMedicalId(c echo.Context) error {
	id := c.Param("medicalID")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	res, err := u.citizenServ.GetSessionByMedicalId(uint(convId))

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
