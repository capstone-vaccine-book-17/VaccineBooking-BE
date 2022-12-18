package adminDto

type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginJWT struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type RegisterAdminDto struct {
	RoleId    uint   `json:"role_id"`
	MedicalId uint   `json:"medical_facilitys_id"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
