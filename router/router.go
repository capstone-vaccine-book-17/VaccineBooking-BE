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

	v1.GET("/", adminController.GetDashboard)

	// TODO AUTH ADMIN
	e.POST("/auth/login", adminController.LoginAdmin)

	// TODO ROLES

	v1_roles := v1.Group("/role")
	v1_roles.POST("/", adminController.CreateRoles, m.Authorization)

	v1_session := v1.Group("/session")
	{
		v1_session.POST("/", adminController.CreateSession)
		v1_session.GET("/", adminController.GetAllSession)
		v1_session.GET("/:id", adminController.GetSessionById)
		v1_session.PUT("/:id", adminController.UpdateSession)
		v1_session.DELETE("/:id", adminController.DeleteSession)
	}

	v1_manageVaccine := v1.Group("/vaccine")
	{
		v1_manageVaccine.POST("/create", adminController.CreateVaccine)
		v1_manageVaccine.GET("/view", adminController.ViewAllVaccine)
		v1_manageVaccine.PUT("/update/:id", adminController.UpdateVaccine)
		v1_manageVaccine.DELETE("/delete/:id", adminController.DeleteVaccine)
	}

	v1_booking := v1.Group("/booking")
	{
		v1_booking.GET("/", adminController.GetAllBooking)
		v1_booking.POST("/", adminController.CreateBooking)
		v1_booking.PUT("/:id", adminController.UpdateBooking)
		v1_booking.GET("/:id", adminController.GetBookingById)
		v1_booking.DELETE("/:id", adminController.DeleteBooking)
	}

}
