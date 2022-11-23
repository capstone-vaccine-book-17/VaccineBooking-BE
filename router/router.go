package router

import (
	"capstone_vaccine/controller/adminController"
	m "capstone_vaccine/middleware"
	"capstone_vaccine/repository/adminRepository"
	"capstone_vaccine/service/adminService"
	"capstone_vaccine/utils"
	"os"

	"github.com/labstack/echo/middleware"

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
	v1.Use(middleware.JWT([]byte(os.Getenv("JWT_KEY"))))

	// TODO AUTH ADMIN
	e.POST("/auth/login", adminController.LoginAdmin)

	// TODO ROLES

	v1_roles := v1.Group("/role")

	v1_roles.POST("/", adminController.CreateRoles)
}
