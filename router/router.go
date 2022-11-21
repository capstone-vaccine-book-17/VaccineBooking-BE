package router

import (
	m "capstone_vaccine/middleware"
	"capstone_vaccine/utils"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	m.LogMiddleware(e)
	e.Validator = &utils.CustomValidator{}
}
