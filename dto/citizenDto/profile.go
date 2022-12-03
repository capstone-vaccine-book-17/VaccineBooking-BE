package citizenDto

type ProfileDTO struct {
	Name  string `json:"name"`
	Nik   string `json:"nik"`
	Image string `json:"image"`
}

type ProfileReq struct {
	CitizenID uint   `json:"citizen_id"`
	Image     string `json:"image"`
	NewEmail  string `json:"new_email"`
}

type PersonalData struct {
	Name    string `json:"name"`
	Nik     string `json:"nik"`
	Age     uint   `json:"age"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Gender  string `json:"gender"`
	Dob     string `json:"dob"`
}

type AddressCitizenReq struct {
	CitizenID  uint   `json:"citizen_id"`
	AddressID  uint   `json:"address_id"`
	NewAddress string `json:"new_address"`
	Province   string `json:"province"`
	City       string `json:"city"`
	PostCode   string `json:"post_code"`
}
type AddressResp struct {
	Address  string `json:"address"`
	Province string `json:"province"`
	City     string `json:"city"`
	PostCode string `json:"post_code"`
}

type UpdateEmail struct {
	CitizenID uint   `json:"citizen_id"`
	Email     string `json:"email" validate:"required,email"`
}

type UpdatePassword struct {
	CitizenID   uint   `json:"citizen_id"`
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required"`
}
