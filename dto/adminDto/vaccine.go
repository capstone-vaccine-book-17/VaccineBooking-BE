package adminDto

type VaccineDTO struct {
	Name    string `json:"name" validate:"required"`
	Kuota   int    `json:"kuota" validate:"required"`
	Expired string `json:"expired" validate:"required"`
}

type VaccineRequest struct {
	MedicalFacilitysId uint   `json:"medical_facilitys_id" validate:"required"`
	Name               string `json:"name" validate:"required"`
	Kuota              int    `json:"kuota" validate:"required"`
	Expired            string `json:"expired" validate:"required"`
}
