package citizenDto

type ProfileDTO struct {
	Name  string `json:"name"`
	Nik   string `json:"nik"`
	Image string `json:"image"`
}

type ProfileReq struct {
	CitizenID uint   `json:"citizen_id"`
	Image     string `json:"image"`
}
