package adminController

import "capstone_vaccine/service/adminService"

type AdminController interface{}

type adminController struct {
	adminServ adminService.AdminService
}

func NewAdminController(adminService adminService.AdminService) *adminController {
	return &adminController{
		adminServ: adminService,
	}
}

// TODO ADMIN CONTROLLER HERE
