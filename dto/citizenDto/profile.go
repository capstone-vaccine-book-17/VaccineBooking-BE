package citizenDto

type ProfileDTO struct {
	Name string `json:"name"`
	Nik  string `json:"nik"`
}

type ProfileReq struct {
	CitizenID uint `json:"citizen_id"`
}
