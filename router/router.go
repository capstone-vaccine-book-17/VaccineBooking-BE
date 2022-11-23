package router

import (
	"capstone_vaccine/controller/adminController"
	m "capstone_vaccine/middleware"
	"capstone_vaccine/repository/adminRepository"
	"capstone_vaccine/service/adminService"
	"capstone_vaccine/utils"

	"github.com/go-playground/validator"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	m.LogMiddleware(e)
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	// TODO REPOSITORY
	adminRepository := adminRepository.NewAdminRepository(db)

	// TODO SERVICE
	adminService := adminService.NewAdminService(adminRepository)

	// TODO CONTROLLER
	adminController := adminController.NewAdminController(adminService)

	// TODO ADMIN ROUTE
	v1 := e.Group("/v1")
	// TODO AUTH ADMIN

	// TODO ROLES
}
