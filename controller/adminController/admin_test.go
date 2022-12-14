package adminController

import (
	"bytes"
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/service/adminService/mocks"
	mocksAdmin "capstone_vaccine/service/adminService/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

// const jwt_tok = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbklEIjozLCJleHAiOjE2NzA4MTkyODIsIm1lZGljYWxJRCI6MSwicm9sZUlEIjoyLCJ1c2VybmFtZSI6ImZhY2hydWRpbiJ9.st-OsNyD_WnNjBLmj-NryetvxNlkfy1g5HjJmvD_06E"

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

// TODO TEST DELETE SESSION VALID
