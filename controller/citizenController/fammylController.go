package citizenController

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/utils"
	"net/http"
	"strconv"

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

// TODO VIEW Member Family
func (u *citizenController) GetFamilys(c echo.Context) error {
	citizenID, _ := middleware.ClaimData(c, "citizenID")
	conv_citizenID := citizenID.(float64)
	conv := uint(conv_citizenID)

	var payloads citizenDto.FamilyReq

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	temp := citizenDto.FamilyReq{
		CitizenId: conv,
	}

	res, err := u.citizenServ.GetFamilys(temp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Family Members",
		Code:    http.StatusOK,
		Data:    res,
	})

}

// TODO Delete Member
func (u *citizenController) DeleteMember(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})

	}

	member := citizenDto.FamilylDTO{
		FamilyId: uint(convId),
	}
	errs := u.citizenServ.DeleteMember(member)

	if errs != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Delete Success",
		Code:    http.StatusOK,
	})

}
