package adminController

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/utils"
	"context"
	"net/http"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// TODO Update Profile
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
	hash, _ := HashPassword(payloads.NewPassword)

	temp := adminDto.ProfileRequest{
		AdminID:            conv_,
		MedicalFacilitysId: conv,
		Name:               payloads.Name,
		Image:              payloads.Image,
		Address:            payloads.Address,
		Username:           payloads.Username,
		NewPassword:        hash,
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

func (u *adminController) UpdateImage(c echo.Context) error {
	cld, _ := cloudinary.NewFromURL("cloudinary://593273685751979:K3Apu1EGSQfIoi9Tzn3zzdGdd6A@dst6d6bj6")
	medicalID, _ := middleware.ClaimData(c, "medicalID")
	conv_medicalID := medicalID.(float64)
	conv := uint(conv_medicalID)

	fileHeader, _ := c.FormFile("image")

	file, _ := fileHeader.Open()

	ctx := context.Background()

	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{})

	if err != nil {
		return err
	}

	var payloads adminDto.ProfileRequest

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	temp := adminDto.ProfileRequest{
		MedicalFacilitysId: conv,
		Image:              result.SecureURL,
	}
	res, err := u.adminServ.UpdateImage(temp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Successfully uploaded the file",
		Code:    http.StatusOK,
		Data:    res.Image,
	})
}
