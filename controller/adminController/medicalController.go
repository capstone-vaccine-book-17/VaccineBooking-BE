package adminController

// func (u *adminController) CreateMedical(c echo.Context) error {
// 	var payloads adminDto.MedicalDto

// 	if err := c.Bind(&payloads); err != nil {
// 		return err
// 	}

// 	if err := c.Validate(payloads); err != nil {
// 		return c.JSON(http.StatusBadRequest, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusBadRequest,
// 		})
// 	}

// 	res, err := u.adminServ.CreateMedical(payloads)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusInternalServerError,
// 		})
// 	}

// 	return c.JSON(http.StatusOK, utils.Response{
// 		Message: "create success",
// 		Code:    http.StatusOK,
// 		Data:    res,
// 	})
// }
