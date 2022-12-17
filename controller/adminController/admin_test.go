package adminController

import (
	"bytes"
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/service/adminService/mocks"
	mocksAdmin "capstone_vaccine/service/adminService/mocks"
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dgrijalva/jwt-go"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

var (
	mockServ   *mocksAdmin.MockAdmin
	controller *adminController
)

func InitEcho() *echo.Echo {
	e := echo.New()

	return e
}

type JwtResponse struct {
	Token string `json:"token"`
}

var jwtToken *jwt.Token

func TestMain(m *testing.M) {
	mocks := &mocks.MockAdmin{}
	mockServ = mocks

	controller = &adminController{
		adminServ: mocks,
	}

	m.Run()
}

// TODO TEST VALID AND INVALID REGISTER ADMIN
func TestRegisterAdmin_Valid(t *testing.T) {
	data := adminDto.RegisterAdminDto{
		RoleId:    1,
		MedicalId: 1,
		Username:  "testing",
		Password:  "testing",
	}
	mockServ.On("RegisterAdmin", data).Return(data, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               adminDto.RegisterAdminDto
		HasReturnBody      bool
		ExpectedBody       adminDto.RegisterAdminDto
	}{
		{
			"success",
			http.StatusOK,
			"POST",
			data,
			true,
			data,
		},
	}
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			r.Header.Add("Content-Type", "application/json")

			e.Validator = &CustomValidator{validator: validator.New()}
			assert.NoError(t, ctx.Validate(v.Body))

			err := controller.RegisterAdmin(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}

				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				data := resp["data"]
				conv, _ := data.(map[string]interface{})

				assert.Equal(t, v.ExpectedBody.Username, conv["username"])
			}
		})
	}
}

func TestRegisterAdmin_InValid(t *testing.T) {
	data := adminDto.RegisterAdminDto{
		RoleId:    1,
		MedicalId: 1,
		Username:  "",
		Password:  "",
	}
	mockServ.On("RegisterAdmin", data).Return(data, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               adminDto.RegisterAdminDto
		HasReturnBody      bool
		ExpectedBody       string
	}{
		{
			"success",
			http.StatusBadRequest,
			"POST",
			data,
			true,
			"Key: 'RegisterAdminDto.Username' Error:Field validation for 'Username' failed on the 'required' tag\nKey: 'RegisterAdminDto.Password' Error:Field validation for 'Password' failed on the 'required' tag",
		},
	}
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			r.Header.Add("Content-Type", "application/json")

			e.Validator = &CustomValidator{validator: validator.New()}
			assert.Equal(t, ctx.Validate(v.Body), ctx.Validate(v.Body))

			err := controller.RegisterAdmin(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}

				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				assert.Equal(t, v.ExpectedBody, resp["message"])
			}
		})
	}
}

// TODO TEST LOGIN VALID AND INVALID
func TestLoginAdmin_Valid(t *testing.T) {
	data := adminDto.LoginDTO{
		Username: "testing",
		Password: "testing",
	}
	mockServ.On("LoginAdmin", data).Return(adminDto.LoginJWT{
		Username: "testing",
		Token:    "blabla",
	}, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               adminDto.LoginDTO
		HasReturnBody      bool
		ExpectedBody       adminDto.LoginJWT
	}{
		{
			"success",
			http.StatusOK,
			"POST",
			data,
			true,
			adminDto.LoginJWT{
				Username: "testing",
				Token:    "blabla",
			},
		},
	}
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			r.Header.Add("Content-Type", "application/json")

			e.Validator = &CustomValidator{validator: validator.New()}
			assert.NoError(t, ctx.Validate(v.Body))

			err := controller.LoginAdmin(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}

				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				data := resp["data"]
				conv, _ := data.(map[string]interface{})

				assert.Equal(t, v.ExpectedBody.Token, conv["token"])

				response := &JwtResponse{}

				jwtToken = &jwt.Token{
					Raw: response.Token,
					Claims: jwt.MapClaims{
						"medicalID": float64(1),
						"adminID":   float64(1),
					},
				}
			}
		})
	}
}

func TestLoginAdmin_InValid_Empty(t *testing.T) {
	data := adminDto.LoginDTO{
		Username: "",
		Password: "testing",
	}
	mockServ.On("LoginAdmin", data).Return(adminDto.LoginJWT{
		Username: "testing",
		Token:    "blabla",
	}, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               adminDto.LoginDTO
		HasReturnBody      bool
		ExpectedBody       string
	}{
		{
			"success",
			http.StatusBadRequest,
			"POST",
			data,
			true,
			"Key: 'LoginDTO.Username' Error:Field validation for 'Username' failed on the 'required' tag",
		},
	}
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			r.Header.Add("Content-Type", "application/json")

			e.Validator = &CustomValidator{validator: validator.New()}
			assert.Equal(t, ctx.Validate(v.Body), ctx.Validate(v.Body))

			err := controller.LoginAdmin(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}

				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				assert.Equal(t, v.ExpectedBody, resp["message"])
			}
		})
	}
}

// TODO TEST GET DASHBOARD VALID
func TestGetDashboard_Valid(t *testing.T) {
	data := adminDto.CountDashboard{
		VaccineAvailable:   "1000",
		BookingToday:       "10",
		BookingsRegistered: "100",
	}
	mockServ.On("GetDashboard", uint(1)).Return(data, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		ExpectedBody       adminDto.CountDashboard
	}{
		{
			"success",
			http.StatusOK,
			"GET",
			data,
		},
	}

	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			ctx.Set("user", jwtToken)

			err := controller.GetDashboard(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			var resp map[string]interface{}

			_ = json.NewDecoder(w.Result().Body).Decode(&resp)

			data := resp["data"]
			conv, _ := data.(map[string]interface{})

			assert.Equal(t, v.ExpectedBody.BookingToday, conv["booking_today"])
		})
	}
}

// TODO CREATE SESSION VALID
func TestCreateSession_Valid(t *testing.T) {
	data := adminDto.SessionRequest{
		Name:               "sesi 1",
		MedicalFacilitysId: 1,
		VaccineId:          1,
		StartTime:          "10:30",
		Kuota:              1000,
		Dosis:              "Pertama",
		EndTime:            "12:30",
		Date:               "2022-10-22",
	}
	mockServ.On("CreateSession", data).Return(adminDto.SessionDTO{
		Name:      "sesi 1",
		VaccineId: 1,
		StartTime: "10:30",
		Kuota:     1000,
		Dosis:     "Pertama",
		EndTime:   "12:30",
		Date:      "2022-10-22",
	}, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               adminDto.SessionRequest
		HasReturnBody      bool
		ExpectedBody       adminDto.SessionDTO
	}{
		{
			"success",
			http.StatusOK,
			"POST",
			data,
			true,
			adminDto.SessionDTO{
				Name:      "sesi 1",
				VaccineId: 1,
				StartTime: "10:30",
				Kuota:     1000,
				Dosis:     "Pertama",
				EndTime:   "12:30",
				Date:      "2022-10-22",
			},
		},
	}
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/v1/session/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			r.Header.Add("Content-Type", "application/json")

			ctx.Set("user", jwtToken)

			err := controller.CreateSession(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}

				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				data := resp["data"]
				conv, _ := data.(map[string]interface{})

				assert.Equal(t, v.ExpectedBody.Name, conv["name"])
			}
		})
	}
}

// TODO TEST GET ALL SESSION VALID
func TestGetAllSession_Valid(t *testing.T) {
	data := []adminDto.SessionWithStatusDTO{
		{
			SessionId:   1,
			Name:        "sesi 1",
			VaccineName: "Moderna",
			StartTime:   "15:03",
			Kuota:       1000,
			Dosis:       "Kedua",
			EndTime:     "16:03",
			Date:        "2022-12-29",
			Status:      "process",
		},
	}
	mockServ.On("GetAllSession", uint(1)).Return(data, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		ExpectedBody       []adminDto.SessionWithStatusDTO
	}{
		{
			"success",
			http.StatusOK,
			"GET",
			data,
		},
	}

	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/v1/session/", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			ctx.Set("user", jwtToken)

			err := controller.GetAllSession(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			var resp map[string][]adminDto.SessionWithStatusDTO

			_ = json.NewDecoder(w.Result().Body).Decode(&resp)

			assert.Equal(t, v.ExpectedBody[0].Name, resp["data"][0].Name)
		})
	}
}

// TODO TEST GET SESSION BY ID VALID
func TestGetSessionById_Valid(t *testing.T) {
	data := adminDto.SessionWithStatusDTO{
		SessionId: 1,
	}
	mockServ.On("GetSessionById", data.SessionId).Return(adminDto.SessionWithStatusDTO{
		SessionId:   1,
		Name:        "sesi 1",
		VaccineName: "Moderna",
		StartTime:   "15:03",
		Kuota:       1000,
		Dosis:       "Kedua",
		EndTime:     "16:03",
		Date:        "2022-12-29",
		Status:      "process",
	}, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               adminDto.SessionWithStatusDTO
		HasReturnBody      bool
		ExpectedBody       adminDto.SessionWithStatusDTO
	}{
		{
			"success",
			http.StatusOK,
			"GET",
			data,
			true,
			adminDto.SessionWithStatusDTO{
				SessionId:   1,
				Name:        "sesi 1",
				VaccineName: "Moderna",
				StartTime:   "15:03",
				Kuota:       1000,
				Dosis:       "Kedua",
				EndTime:     "16:03",
				Date:        "2022-12-29",
				Status:      "process",
			},
		},
	}

	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/v1/session/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := controller.GetSessionById(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}

				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				data := resp["data"]
				conv, _ := data.(map[string]interface{})

				assert.Equal(t, v.ExpectedBody.Name, conv["name"])
			}
		})
	}
}

// TODO TEST UPDATE SESSION VALID AND INVALID
func TestUpdateSession_Valid(t *testing.T) {
	data := adminDto.SessionRequestUpdate{
		SessionId: 1,
		Name:      "sesi 1 edited",
		VaccineId: 1,
		StartTime: "10:15 edited",
		Kuota:     100,
		Dosis:     "Kedua",
		EndTime:   "11:15 edited",
	}
	mockServ.On("UpdateSession", data).Return(data, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               adminDto.SessionRequestUpdate
		HasReturnBody      bool
		ExpectedBody       adminDto.SessionRequestUpdate
	}{
		{
			"success",
			http.StatusOK,
			"POST",
			data,
			true,
			data,
		},
	}
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/v1/session/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			r.Header.Add("Content-Type", "application/json")

			e.Validator = &CustomValidator{validator: validator.New()}
			assert.NoError(t, ctx.Validate(v.Body))

			err := controller.UpdateSession(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}

				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				data := resp["data"]
				conv, _ := data.(map[string]interface{})

				assert.Equal(t, v.ExpectedBody.Name, conv["name"])
			}
		})
	}
}
func TestUpdateSession_InValid(t *testing.T) {
	data := adminDto.SessionRequestUpdate{
		SessionId: 1,
		Name:      "sesi 1 edited",
		VaccineId: 1,
		StartTime: "10:15 edited",
		Kuota:     100,
		Dosis:     "",
		EndTime:   "11:15 edited",
	}
	mockServ.On("UpdateSession", data).Return(data, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               adminDto.SessionRequestUpdate
		HasReturnBody      bool
		ExpectedBody       string
	}{
		{
			"success",
			http.StatusBadRequest,
			"POST",
			data,
			true,
			"Key: 'SessionRequestUpdate.Dosis' Error:Field validation for 'Dosis' failed on the 'required' tag",
		},
	}
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/v1/session/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			r.Header.Add("Content-Type", "application/json")

			e.Validator = &CustomValidator{validator: validator.New()}
			assert.Equal(t, ctx.Validate(v.Body), ctx.Validate(v.Body))

			err := controller.UpdateSession(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}

				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				assert.Equal(t, v.ExpectedBody, resp["message"])
			}
		})
	}
}

// TODO TEST DELETE SESSION VALID AND INVALID
func TestDeleteSession_Valid(t *testing.T) {
	data := adminDto.SessionWithStatusDTO{
		SessionId: 1,
	}

	mockServ.On("DeleteSession", data.SessionId).Return(nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		ExpectedBody       string
	}{
		{
			"success",
			http.StatusOK,
			"DELETE",
			"success",
		},
	}

	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/v1/session/1", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/v1/session/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			ctx.Set("user", jwtToken)

			err := controller.DeleteSession(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			var resp map[string]interface{}

			_ = json.NewDecoder(w.Result().Body).Decode(&resp)

			assert.Equal(t, v.ExpectedBody, resp["message"])
		})
	}
}

func TestDeleteSession_InValid(t *testing.T) {
	data := adminDto.SessionWithStatusDTO{
		SessionId: 1,
	}

	mockServ.On("DeleteSession", data.SessionId).Return(nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		ExpectedBody       string
	}{
		{
			"success",
			http.StatusBadRequest,
			"DELETE",
			"strconv.Atoi: parsing \"a\": invalid syntax",
		},
	}

	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/v1/session/a", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/v1/session/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("a")

			ctx.Set("user", jwtToken)

			err := controller.DeleteSession(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			var resp map[string]interface{}

			_ = json.NewDecoder(w.Result().Body).Decode(&resp)

			assert.Equal(t, v.ExpectedBody, resp["message"])
		})
	}
}

// TODO TEST CREATE VACCINE VALID And INVALID
func TestCreateVaccine_Valid(t *testing.T) {
	data := adminDto.VaccineRequest{
		Name:               "moderna",
		MedicalFacilitysId: 1,
		Kuota:              1000,
		Expired:            "2022-12-20",
	}
	mockServ.On("CreateVaccine", data).Return(adminDto.VaccineResponse{
		Name:    "moderna",
		Kuota:   1000,
		Expired: "2022-12-20",
	}, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               adminDto.VaccineRequest
		HasReturnBody      bool
		ExpectedBody       adminDto.VaccineResponse
	}{
		{
			"success",
			http.StatusOK,
			"POST",
			data,
			true,
			adminDto.VaccineResponse{
				Name:    "moderna",
				Kuota:   1000,
				Expired: "2022-12-20",
			},
		},
	}
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/v1/vaccine/create", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			r.Header.Add("Content-Type", "application/json")
			e.Validator = &CustomValidator{validator: validator.New()}
			assert.Equal(t, ctx.Validate(v.Body), ctx.Validate(v.Body))

			ctx.Set("user", jwtToken)

			err := controller.CreateVaccine(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}

				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				data := resp["data"]
				conv, _ := data.(map[string]interface{})

				assert.Equal(t, v.ExpectedBody.Name, conv["name"])
			}
		})
	}
}

func TestCreateVaccine_InValid(t *testing.T) {
	data := adminDto.VaccineRequest{
		Name:               "",
		MedicalFacilitysId: 1,
		Kuota:              1000,
		Expired:            "2022-12-20",
	}
	mockServ.On("CreateVaccine", data).Return(adminDto.VaccineResponse{}, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               adminDto.VaccineRequest
		HasReturnBody      bool
		ExpectedBody       string
	}{
		{
			"success",
			http.StatusBadRequest,
			"POST",
			data,
			true,
			"Key: 'VaccineRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag",
		},
	}
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/v1/vaccine/create", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			r.Header.Add("Content-Type", "application/json")
			e.Validator = &CustomValidator{validator: validator.New()}
			assert.Equal(t, ctx.Validate(v.Body), ctx.Validate(v.Body))

			ctx.Set("user", jwtToken)

			err := controller.CreateVaccine(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}

				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				assert.Equal(t, v.ExpectedBody, resp["message"])
			}
		})
	}
}

// TODO TEST VIEW ALL VACCINE
func TestViewAllVaccine_Valid(t *testing.T) {
	data := []adminDto.VaccineDTO{
		{
			VaccineID: 1,
			Name:      "astra",
			Kuota:     1000,
			Expired:   "2023-12-2",
		},
	}
	mockServ.On("ViewAllVaccine", uint(1)).Return(data, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		ExpectedBody       []adminDto.VaccineDTO
	}{
		{
			"success",
			http.StatusOK,
			"GET",
			data,
		},
	}

	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/v1/vaccine/view", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			ctx.Set("user", jwtToken)

			err := controller.ViewAllVaccine(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			var resp map[string][]adminDto.VaccineDTO

			_ = json.NewDecoder(w.Result().Body).Decode(&resp)

			assert.Equal(t, v.ExpectedBody[0].Name, resp["data"][0].Name)
		})
	}
}

// TODO TEST UPDATE Vaccine VALID AND INVALID
func TestUpdateVaccine_Valid(t *testing.T) {
	data := adminDto.VaccineDTO{
		VaccineID: 1,
		Name:      "astra edited",
		Kuota:     1000,
		Expired:   "2024-12-1 edited",
	}
	mockServ.On("UpdateVaccine", data).Return(data, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               adminDto.VaccineDTO
		HasReturnBody      bool
		ExpectedBody       adminDto.VaccineDTO
	}{
		{
			"success",
			http.StatusOK,
			"POST",
			data,
			true,
			data,
		},
	}
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Set("user", jwtToken)
			ctx.SetPath("/v1/vaccine/update/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			r.Header.Add("Content-Type", "application/json")

			e.Validator = &CustomValidator{validator: validator.New()}
			assert.NoError(t, ctx.Validate(v.Body))

			err := controller.UpdateVaccine(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}

				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				data := resp["data"]
				conv, _ := data.(map[string]interface{})

				assert.Equal(t, v.ExpectedBody.Name, conv["name"])
			}
		})
	}
}

func TestUpdateVaccine_InValid(t *testing.T) {
	data := adminDto.VaccineDTO{
		VaccineID: 1,
		Name:      "",
		Kuota:     1000,
		Expired:   "2024-12-1 edited",
	}
	mockServ.On("UpdateVaccine", data).Return(data, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               adminDto.VaccineDTO
		HasReturnBody      bool
		ExpectedBody       string
	}{
		{
			"success",
			http.StatusBadRequest,
			"POST",
			data,
			true,
			"Key: 'VaccineDTO.Name' Error:Field validation for 'Name' failed on the 'required' tag",
		},
	}
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Set("user", jwtToken)
			ctx.SetPath("/v1/session/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			r.Header.Add("Content-Type", "application/json")

			e.Validator = &CustomValidator{validator: validator.New()}
			assert.Equal(t, ctx.Validate(v.Body), ctx.Validate(v.Body))

			err := controller.UpdateVaccine(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}

				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				assert.Equal(t, v.ExpectedBody, resp["message"])
			}
		})
	}
}

// TODO TEST DELETE Vaccine VALID AND INVALID
func TestDeleteVaccine_Valid(t *testing.T) {
	data := adminDto.VaccineDTO{
		VaccineID: 1,
	}

	mockServ.On("DeleteVaccine", data.VaccineID, uint(1)).Return(nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		ExpectedBody       string
	}{
		{
			"success",
			http.StatusOK,
			"DELETE",
			"Vaccine Berhasil Dihapus",
		},
	}

	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/v1/vaccine/delete/1", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/v1/vaccine/delete/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			ctx.Set("user", jwtToken)

			err := controller.DeleteVaccine(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			var resp map[string]interface{}

			_ = json.NewDecoder(w.Result().Body).Decode(&resp)

			assert.Equal(t, v.ExpectedBody, resp["message"])
		})
	}
}

func TestDeleteVaccine_InValid(t *testing.T) {
	data := adminDto.VaccineDTO{
		VaccineID: 1,
	}

	mockServ.On("DeleteVaccine", data.VaccineID, uint(1)).Return(nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		ExpectedBody       string
	}{
		{
			"success",
			http.StatusBadRequest,
			"DELETE",
			"strconv.Atoi: parsing \"a\": invalid syntax",
		},
	}

	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/v1/vaccine/delete/a", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/v1/vaccine/delete/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("a")

			ctx.Set("user", jwtToken)

			err := controller.DeleteVaccine(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			var resp map[string]interface{}

			_ = json.NewDecoder(w.Result().Body).Decode(&resp)

			assert.Equal(t, v.ExpectedBody, resp["message"])
		})
	}
}

// TODO TEST GET BY ID VACCINE VALID And INVALID
func TestGetVaccineById_Valid(t *testing.T) {
	mockServ.On("GetVaccineById", uint(1), uint(1)).Return(adminDto.VaccineDTO{
		VaccineID: 1,
		Name:      "astra",
		Kuota:     1000,
		Expired:   "2023-12-2",
	}, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               uint
		HasReturnBody      bool
		ExpectedBody       adminDto.VaccineDTO
	}{
		{
			"success",
			http.StatusOK,
			"GET",
			1,
			true,
			adminDto.VaccineDTO{
				VaccineID: 1,
				Name:      "astra",
				Kuota:     1000,
				Expired:   "2023-12-2",
			},
		},
	}

	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Set("user", jwtToken)
			ctx.SetPath("/v1/vaccine/view/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := controller.GetVaccineById(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			if v.HasReturnBody {
				var resp map[string]interface{}

				_ = json.NewDecoder(w.Result().Body).Decode(&resp)

				data := resp["data"]
				conv, _ := data.(map[string]interface{})

				assert.Equal(t, v.ExpectedBody.Name, conv["name"])
			}
		})
	}
}

func TestGetVaccineById_InValid(t *testing.T) {

	mockServ.On("GetVaccineById", uint(1), uint(1)).Return(adminDto.VaccineDTO{
		VaccineID: 1,
		Name:      "astra",
		Kuota:     1000,
		Expired:   "2023-12-2",
	}, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Body               uint
		Method             string
		ExpectedBody       string
	}{
		{
			"success",
			http.StatusBadRequest,
			1,
			"GET",
			"strconv.Atoi: parsing \"a\": invalid syntax",
		},
	}

	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/v1/vaccine/view/a", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/v1/vaccine/view/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues("a")

			ctx.Set("user", jwtToken)

			err := controller.GetVaccineById(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			var resp map[string]interface{}

			_ = json.NewDecoder(w.Result().Body).Decode(&resp)

			assert.Equal(t, v.ExpectedBody, resp["message"])
		})
	}
}

// TODO TEST GET PROFILE
func TestGetProfile_Valid(t *testing.T) {
	data := []adminDto.ProfilDTO{
		{
			Name:              "dika",
			Image:             "url",
			Address:           "jalan",
			ResponsiblePerson: "dika",
			Username:          "dika",
		},
	}
	payload := adminDto.ProfileRequest{
		AdminID: 1,
	}
	mockServ.On("GetProfile", payload).Return(data, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		ExpectedBody       []adminDto.ProfilDTO
	}{
		{
			"success",
			http.StatusOK,
			"GET",
			data,
		},
	}

	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			r := httptest.NewRequest(v.Method, "/v1/profile/", nil)
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)

			ctx.Set("user", jwtToken)

			err := controller.GetProfile(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

			var resp map[string][]adminDto.ProfilDTO

			_ = json.NewDecoder(w.Result().Body).Decode(&resp)

			assert.Equal(t, v.ExpectedBody[0].Name, resp["data"][0].Name)
		})
	}
}

// UPDATE PROFILE
func TestUpdateProfile_Valid(t *testing.T) {
	data := adminDto.ProfileRequest{
		AdminID:            1,
		MedicalFacilitysId: 1,
		Name:               "dika",
		Address:            "jalan",
		ResponsiblePerson:  "dika",
		Username:           "dika",
		NewPassword:        "rama",
		Password:           "dika",
	}
	mockServ.On("UpdateProfile", data).Return(data, nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               adminDto.ProfileRequest
		HasReturnBody      bool
	}{
		{
			"success",
			http.StatusOK,
			"PUT",
			data,
			false,
		},
	}
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/v1/profile/update", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.Set("user", jwtToken)

			r.Header.Add("Content-Type", "application/json")

			err := controller.UpdateProfile(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)

		})
	}
}

// TEST UPDATE IMAGE
func TestUpdateImage_Valid(t *testing.T) {
	os.Setenv("CLOUDINARY_URL", "cloudinary://593273685751979:K3Apu1EGSQfIoi9Tzn3zzdGdd6A@dst6d6bj6")
	data := adminDto.ProfileRequest{
		MedicalFacilitysId: 1,
	}

	mockServ.On("UpdateImage", data).Return(nil).Once()

	testCases := []struct {
		Name               string
		ExpectedStatusCode int
		Method             string
		Body               adminDto.ProfileRequest
		HasReturnBody      bool
	}{
		{
			"success",
			http.StatusOK,
			"PUT",
			data,
			false,
		},
	}
	for _, v := range testCases {
		t.Run(v.Name, func(t *testing.T) {
			// setup file
			pr, pw := io.Pipe()
			writer := multipart.NewWriter(pw)

			go func() {
				defer writer.Close()

				part, err := writer.CreateFormFile("image", "image.png")
				if err != nil {
					t.Error(err)
				}

				img := createImage()

				err = png.Encode(part, img)
				if err != nil {
					t.Error(err)
				}
			}()

			r := httptest.NewRequest(v.Method, "/v1/profile/image", pr)
			w := httptest.NewRecorder()

			r.Header.Add("Content-Type", writer.FormDataContentType())

			e := echo.New()
			ctx := e.NewContext(r, w)

			ctx.Set("user", jwtToken)

			err := controller.UpdateImage(ctx)
			assert.NoError(t, err)

			assert.Equal(t, v.ExpectedStatusCode, w.Result().StatusCode)
		})
	}
}

func createImage() *image.RGBA {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}

	// Encode as PNG.
	os.Create("image.png")

	return img
}
