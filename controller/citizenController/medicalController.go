package citizenController

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// TODO GET ALL MEDICAL BY CITY AND BY SEARCH
func (u *citizenController) GetMedicalByCity(c echo.Context) error {

	citizenID, _ := middleware.ClaimData(c, "citizenID")
	conv_citizenID := citizenID.(float64)
	citizen_id := uint(conv_citizenID)

	var payloads citizenDto.SearchKey

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	temp := citizenDto.SearchKey{
		CitizenId: citizen_id,
		S:         payloads.S,
		Q:         payloads.Q,
	}
	res, err := u.citizenServ.GetMedicalByCity(temp)

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

// TODO GET MEDICAL BY ID
func (u *citizenController) GetMedicalById(c echo.Context) error {

	id := c.Param("medicalID")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	res, err := u.citizenServ.GetMedicalById(uint(convId))

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
