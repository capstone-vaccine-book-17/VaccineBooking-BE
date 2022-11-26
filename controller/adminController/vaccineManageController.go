package adminController

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/utils"
	"net/http"

	"github.com/labstack/echo"
)

func (u *adminController) CreateVaccine(c echo.Context) error {
	medicalID, _ := middleware.ClaimData(c, "medicalID")

	conv_medicalID := medicalID.(float64)

	conv := uint(conv_medicalID)

	var payloads adminDto.VaccineRequest

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	temp := adminDto.VaccineRequest{
		MedicalFacilitysId: conv,
		Name:               payloads.Name,
		Kuota:              payloads.Kuota,
		Expired:            payloads.Expired,
	}

	res, err := u.adminServ.CreateVaccine(temp)
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
