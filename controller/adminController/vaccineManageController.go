package adminController

// // TODO Create Vaccine
// func (u *adminController) CreateVaccine(c echo.Context) error {
// 	medicalID, _ := middleware.ClaimData(c, "medicalID")

// 	conv_medicalID := medicalID.(float64)

// 	conv := uint(conv_medicalID)

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
// 	temp := adminDto.VaccineRequest{
// 		MedicalFacilitysId: conv,
// 		Name:               payloads.Name,
// 		Kuota:              payloads.Kuota,
// 		Expired:            payloads.Expired,
// 	}

// 	res, err := u.adminServ.CreateVaccine(temp)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusInternalServerError,
// 		})
// 	}

// 	return c.JSON(http.StatusOK, utils.Response{
// 		Message: "Vaccine Berhasil Ditambahkan",
// 		Code:    http.StatusOK,
// 		Data:    res,
// 	})
// }

// // TODO View All Vaccine
// func (u *adminController) ViewAllVaccine(c echo.Context) error {

// 	res, err := u.adminServ.ViewAllVaccine()
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusBadRequest,
// 		})
// 	}

<<<<<<< HEAD
// 	return c.JSON(http.StatusOK, utils.Response{
// 		Message: "Daftar Seluruh Vaccine",
// 		Code:    http.StatusOK,
// 		Data:    res,
// 	})
// }

// // TODO Update Vaccine
// func (u *adminController) UpdateVaccine(c echo.Context) error {
// 	id := c.Param("id")
// 	convId, err := strconv.Atoi(id)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusBadRequest,
// 		})
// 	}
=======
// TODO View All Vaccine
func (u *adminController) ViewAllVaccine(c echo.Context) error {
	medicalID, _ := middleware.ClaimData(c, "medicalID")
	conv_medicalID := medicalID.(float64)
	conv := uint(conv_medicalID)

	res, err := u.adminServ.ViewAllVaccine(conv)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}
>>>>>>> 5f7d23e513a40b96d601b5af6eb5e857711639d2

// 	var payloads adminDto.VaccineDTO

<<<<<<< HEAD
// 	if err := c.Bind(&payloads); err != nil {
// 		return err
// 	}
=======
// TODO Update Vaccine
func (u *adminController) UpdateVaccine(c echo.Context) error {
	medicalID, _ := middleware.ClaimData(c, "medicalID")
	conv_medicalID := medicalID.(float64)
	conv := uint(conv_medicalID)

	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}
>>>>>>> 5f7d23e513a40b96d601b5af6eb5e857711639d2

// 	if err := c.Validate(&payloads); err != nil {
// 		return c.JSON(http.StatusBadRequest, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusBadRequest,
// 		})
// 	}

// 	Vaccine := adminDto.VaccineDTO{
// 		VaccineID: uint(convId),
// 		Name:      payloads.Name,
// 		Kuota:     payloads.Kuota,
// 		Expired:   payloads.Expired,
// 	}

// 	res, err := u.adminServ.UpdateVaccine(Vaccine)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusInternalServerError,
// 		})
// 	}

<<<<<<< HEAD
// 	return c.JSON(http.StatusOK, utils.Response{
// 		Message: "Update Berhasil Dilakukan",
// 		Code:    http.StatusOK,
// 		Data:    res,
// 	})
// }
=======
	res, err := u.adminServ.UpdateVaccine(Vaccine,conv)
>>>>>>> 5f7d23e513a40b96d601b5af6eb5e857711639d2

// // TODO Delete Vaccine
// func (u *adminController) DeleteVaccine(c echo.Context) error {
// 	id := c.Param("id")
// 	convId, err := strconv.Atoi(id)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusBadRequest,
// 		})
// 	}
// 	vaccine := adminDto.VaccineDTO{
// 		VaccineID: uint(convId),
// 	}

// 	err = u.adminServ.DeleteVaccine(vaccine)

<<<<<<< HEAD
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusBadRequest,
// 		})
// 	}

// 	return c.JSON(http.StatusOK, utils.Response{
// 		Message: "Vaccine Berhasil Dihapus",
// 		Code:    http.StatusOK,
// 	})
// }
=======
// TODO Delete Vaccine
func (u *adminController) DeleteVaccine(c echo.Context) error {
	medicalID, _ := middleware.ClaimData(c, "medicalID")
	conv_medicalID := medicalID.(float64)
	conv := uint(conv_medicalID)

	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}
	vaccine := adminDto.VaccineDTO{
		VaccineID: uint(convId),
	}

	err = u.adminServ.DeleteVaccine(vaccine,conv)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Vaccine Berhasil Dihapus",
		Code:    http.StatusOK,
	})
}

// TODO Get Vaccine
func (u *adminController) GetVaccineById(c echo.Context) error {
	medicalID, _ := middleware.ClaimData(c, "medicalID")
	conv_medicalID := medicalID.(float64)
	conv := uint(conv_medicalID)

	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	res, err := u.adminServ.GetVaccineById(uint(convId), conv)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "Vaccine",
		Code:    http.StatusOK,
		Data:    res,
	})
}
>>>>>>> 5f7d23e513a40b96d601b5af6eb5e857711639d2
