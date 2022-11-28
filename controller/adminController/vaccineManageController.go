package adminController

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/utils"
	"net/http"
	"strconv"

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

func (u *adminController) ViewAllVaccine(c echo.Context) error {

	res, err := u.adminServ.ViewAllVaccine()
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "success",
		Code:    http.StatusOK,
		Data:    res,
	})
}

func (u *adminController) GetVaccineById(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	session := adminDto.VaccineRequest{
		VaccineID: uint(convId),
	}

	res, err := u.adminServ.GetVaccineById(session)

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

func (u *adminController) UpdateVaccine(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	var payloads adminDto.VaccineRequest

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	if err := c.Validate(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	Vaccine := adminDto.VaccineRequest{
		VaccineID: uint(convId),
		Name:      payloads.Name,
		Kuota:     payloads.Kuota,
		Expired:   payloads.Expired,
	}

	res, err := u.adminServ.UpdateVaccine(Vaccine)

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

// func (u *adminController) UpdateVaccine(c echo.Context) error {
// 	id := c.Param("id")
// 	convId, err := strconv.Atoi(id)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusBadRequest,
// 		})
// 	}

// 	var payloads adminDto.VaccineRequest

// 	if err := c.Bind(&payloads); err != nil {
// 		return err
// 	}

// 	if err := c.Validate(&payloads); err != nil {
// 		return c.JSON(http.StatusBadRequest, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusBadRequest,
// 		})
// 	}

// 	session := adminDto.VaccineRequest{
// 		VaccineID:          uint(convId),
// 		Name:               payloads.Name,
// 		Kuota:              payloads.Kuota,
// 		Expired:            payloads.Expired,
// 		MedicalFacilitysId: payloads.MedicalFacilitysId,
// 	}

// 	res, err := u.adminServ.UpdateVaccine(session)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusInternalServerError,
// 		})
// 	}

// 	return c.JSON(http.StatusOK, utils.Response{
// 		Message: "success",
// 		Code:    http.StatusOK,
// 		Data:    res,
// 	})
// }

func (u *adminController) DeleteVaccine(c echo.Context) error {
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}
	vaccine := adminDto.VaccineRequest{
		VaccineID: uint(convId),
	}

	res, err := u.adminServ.DeleteVaccine(vaccine)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "success",
		Code:    http.StatusOK,
		Data:    res,
	})
}
