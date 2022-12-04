package citizenDto

type FammylDTO struct {
	FamilyId uint   `json:"family_id"`
	Relation string `json:"relation"`
	Name     string `json:"name"`
	Nik      string `json:"nik"`
	Age      uint   `json:"age"`
	Gender   string `json:"gender"`
}

type FamilyReq struct {
	CitizenId uint   `json:"citizen_id"`
	Relation  string `json:"relation" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Nik       string `json:"nik" validate:"required"`
	Age       uint   `json:"age" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
}
