package adminDto

type ProfilDTO struct {
	Name     string `json:"name"`
	Image    string `json:"image"`
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ProfileRequest struct {
	AdminID            uint   `json:"admin_id"`
	MedicalFacilitysId uint   `json:"medical_facilitys_id"`
	Name               string `json:"name"`
	Image              string `json:"image"`
	Address            string `json:"address"`
	Username           string `json:"username"`
	Password           string `json:"password"`
}
type Address struct {
	AddressID string `json:"address_id"`
}
