package citizenController

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/utils"
	"net/http"

	"github.com/labstack/echo"
)

func (u *citizenController) CreateFamilyMember(c echo.Context) error {
	citizenID, _ := middleware.ClaimData(c, "citizenID")
	conv_citizenID := citizenID.(float64)
	conv := uint(conv_citizenID)

	var payloads citizenDto.FamilyReq

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	temp := citizenDto.FamilyReq{
		CitizenId: conv,
		Relation:  payloads.Relation,
		Name:      payloads.Name,
		Nik:       payloads.Nik,
		Age:       payloads.Age,
		Gender:    payloads.Gender,
	}

	err := u.citizenServ.CreateFamilyMember(temp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Create new family member Success",
		Code:    http.StatusOK,
	})

}
