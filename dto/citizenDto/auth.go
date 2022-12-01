package citizenDto

type RegisterDto struct {
	Name     string `json:"name" validate:"required"`
	Nik      string `json:"nik" validate:"required,max=16"`
	Address  string `json:"address" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
}

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginJWT struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
