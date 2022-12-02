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

func (u *citizenController) UploadImage(c echo.Context) error {

	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))

	citizenID, _ := middleware.ClaimData(c, "citizenID")

	conv_citizenID := citizenID.(float64)

	conv := uint(conv_citizenID)
	

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
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "MY PROFILE",
		Code:    http.StatusOK,
		Data:    res.Image,
	})

}
