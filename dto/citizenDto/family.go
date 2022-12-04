package citizenDto

type fammylDTO struct{
	FamilyID   uint      `json:"family_id"`
	Relation string      `json:"relation"`
	Name       string    `json:"name"`
	Nik        string    `json:"nik"`
	Age        uint      `json:"age"`
	Gender     string    `json:"gender"`
}

type 