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
	FamilyId  uint   `json:"family_id"`
	Relation  string `json:"relation"`
	Name      string `json:"name"`
	Nik       string `json:"nik"`
	Age       uint   `json:"age"`
	Gender    string `json:"gender"`
}
