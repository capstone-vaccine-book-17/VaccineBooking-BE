package adminController

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/utils"
	"net/http"

	"github.com/labstack/echo"
)

// TODO Get Profile
func (u *adminController) GetProfile(c echo.Context) error {

	adminID, _ := middleware.ClaimData(c, "adminID")
	conv_adminID := adminID.(float64)
	conv_ := uint(conv_adminID)

	var payloads adminDto.ProfileRequest

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	temp := adminDto.ProfileRequest{
		AdminID: conv_,
	}

	res, err := u.adminServ.GetProfile(temp)
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

// TODO Update Profile & Change Password
func (u *adminController) UpdateProfile(c echo.Context) error {

	adminID, _ := middleware.ClaimData(c, "adminID")
	conv_adminID := adminID.(float64)
	conv_ := uint(conv_adminID)

	medicalID, _ := middleware.ClaimData(c, "medicalID")
	conv_medicalID := medicalID.(float64)
	conv := uint(conv_medicalID)

	var payloads adminDto.ProfileRequest

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	temp := adminDto.ProfileRequest{
		AdminID:            conv_,
		MedicalFacilitysId: conv,
		Name:               payloads.Name,
		Image:              payloads.Image,
		Address:            payloads.Address,
		ResponsiblePerson:  payloads.ResponsiblePerson,
		Username:           payloads.Username,
		NewPassword:        payloads.NewPassword,
		Password:           payloads.Password,
	}

	_, err := u.adminServ.UpdateProfile(temp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Update Profile Berhasil dilakukan",
		Code:    http.StatusOK,
	})
}

// Upload Image
func (u *adminController) UpdateImage(c echo.Context) error {

	medicalID, _ := middleware.ClaimData(c, "medicalID")
	conv_medicalID := medicalID.(float64)
	conv := uint(conv_medicalID)
	fileHeader, _ := c.FormFile("image")
	file, _ := fileHeader.Open()

	var payloads adminDto.ProfileRequest

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	temp := adminDto.ProfileRequest{
		MedicalFacilitysId: conv,
	}
	err := u.adminServ.UpdateImage(temp, file)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Successfully uploaded the file",
		Code:    http.StatusOK,
	})
}
