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
