package citizenController

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/middleware"
	"capstone_vaccine/utils"
	"context"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
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

// TODO Update & Upload Image
func (u *citizenController) UploadImage(c echo.Context) error {

	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	medicalID, _ := middleware.ClaimData(c, "citizenID")
	conv_medicalID := medicalID.(float64)
	conv := uint(conv_medicalID)

	fileHeader, _ := c.FormFile("image")

	file, _ := fileHeader.Open()

	ctx := context.Background()

	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{})

	if err != nil {
		return err
	}

	var payloads citizenDto.ProfileReq

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	temp := citizenDto.ProfileReq{
		CitizenID: conv,
		Image:     result.SecureURL,
	}
	res, err := u.citizenServ.UploadImage(temp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
			Data:    temp,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Successfully uploaded the file",
		Code:    http.StatusOK,
		Data:    res.Image,
	})
}

// TODO GET Personal Data
func (u *citizenController) GetPersonalData(c echo.Context) error {
	citizenID, _ := middleware.ClaimData(c, "citizenID")
	conv_citizenID := citizenID.(float64)
	conv := uint(conv_citizenID)

	temp := citizenDto.ProfileReq{
		CitizenID: conv,
	}

	res, err := u.citizenServ.GetPersonalData(temp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Personal Data",
		Code:    http.StatusOK,
		Data:    res,
	})
}

// TODO GET Address Citizen
func (u *citizenController) GetAddress(c echo.Context) error {
	citizenID, _ := middleware.ClaimData(c, "citizenID")
	conv_citizenID := citizenID.(float64)
	conv := uint(conv_citizenID)

	temp := citizenDto.ProfileReq{
		CitizenID: conv,
	}

	res, err := u.citizenServ.GetAddress(temp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Success",
		Code:    http.StatusOK,
		Data:    res,
	})
}

// TODO Update Address
func (u *citizenController) UpdateAddress(c echo.Context) error {
	citizenID, _ := middleware.ClaimData(c, "citizenID")
	conv_citizenID := citizenID.(float64)
	conv := uint(conv_citizenID)

	var payloads citizenDto.AddressCitizenReq

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	temp := citizenDto.AddressCitizenReq{
		CitizenID:  conv,
		NewAddress: payloads.NewAddress,
		Province:   payloads.Province,
		City:       payloads.City,
		PostCode:   payloads.PostCode,
	}

	err := u.citizenServ.UpdateAddress(temp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Update Success",
		Code:    http.StatusOK,
	})
}

// Get Email
func (u *citizenController) GetEmail(c echo.Context) error {
	citizenID, _ := middleware.ClaimData(c, "citizenID")
	conv_citizenID := citizenID.(float64)
	conv := uint(conv_citizenID)

	temp := citizenDto.ProfileReq{
		CitizenID: conv,
	}

	res, err := u.citizenServ.GetEmail(temp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Email",
		Code:    http.StatusOK,
		Data:    res.Email,
	})
}

// TODO UPDATE EMAIL
func (u *citizenController) UpdateEmail(c echo.Context) error {
	citizenID, _ := middleware.ClaimData(c, "citizenID")
	conv_citizenID := citizenID.(float64)
	conv := uint(conv_citizenID)

	var payloads citizenDto.UpdateEmail

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: "format email incorrect",
			Code:    http.StatusBadRequest,
		})
	}

	temp := citizenDto.UpdateEmail{
		CitizenID: conv,
		Email:     payloads.Email,
	}

	err := u.citizenServ.UpdateEmail(temp)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "e-mail changed successfully",
		Code:    http.StatusOK,
	})
}

// TODO UPDATE Password
func (u *citizenController) UpdatePassword(c echo.Context) error {
	citizenID, _ := middleware.ClaimData(c, "citizenID")
	conv_citizenID := citizenID.(float64)
	conv := uint(conv_citizenID)

	var payloads citizenDto.UpdatePassword

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	temp1 := citizenDto.ProfileReq{
		CitizenID: conv,
	}

	check, _ := u.citizenServ.GetEmail(temp1)

	err := utils.CompareHash(check.Password, payloads.OldPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: "password incorrect",
			Code:    http.StatusBadRequest,
		})
	}

	old_pw, _ := u.citizenServ.GetEmail(temp1)
	err = utils.CompareHash(old_pw.Password, payloads.NewPassword)
	if err == nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: "New password is the same as old password",
			Code:    http.StatusBadRequest,
		})
	}

	hash, _ := utils.HashBcrypt(payloads.NewPassword)

	err = utils.CompareHash(hash, payloads.ConfirmNewPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: "confirmation New password is wrong",
			Code:    http.StatusBadRequest,
		})
	}

	temp := citizenDto.UpdatePassword{
		CitizenID:   conv,
		OldPassword: payloads.OldPassword,
		NewPassword: hash,
	}

	_, errs := u.citizenServ.UpdatePassword(temp)

	if errs != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: errs.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "changed successfully",
		Code:    http.StatusOK,
	})
}
