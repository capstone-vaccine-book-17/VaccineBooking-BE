package adminDto

type RoleDTO struct {
	Name string `json:"name" validate:"required"`
}

type RoleResponse struct {
	Name      string `json:"email"`
	CreatedAT string `json:"created_at"`
}
