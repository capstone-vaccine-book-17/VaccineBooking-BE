package adminDto

type VaccineDTO struct {
	VaccineID uint   `json:"vaccine_id"`
	Name      string `json:"name" validate:"required"`
	Kuota     int    `json:"kuota" validate:"required"`
	Expired   string `json:"expired" validate:"required"`
}

type VaccineRequest struct {
	VaccineID          uint   `json:"vaccine_id"`
	MedicalFacilitysId uint   `json:"medical_facilitys_id"`
	Name               string `json:"name" validate:"required"`
	Kuota              int    `json:"kuota" validate:"required"`
	Expired            string `json:"expired" validate:"required"`
}
