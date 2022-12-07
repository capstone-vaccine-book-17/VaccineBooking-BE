package citizenDto

type MedicalDto struct {
	MedicalFacilitysID uint   `json:"medical_facilitys_id"`
	Name               string `json:"name"`
	Address            string `json:"address"`
	Province           string `json:"province"`
	PostCode           string `json:"post_code"`
	Country            string `json:"country"`
	City               string `json:"city"`
}
type SearchKey struct {
	CitizenId uint   `json:"citizen_id"`
	City      string `json:"city"`
	S         string `param:"s" query:"s" json:"s" form:"s"`
	Q         string `param:"q" query:"q" json:"q" form:"q"`
}

type SearchDto struct {
	MedicalFacilitysID uint   `json:"medical_facilitys_id"`
	Name               string `json:"name"`
	Address            string `json:"address"`
	Province           string `json:"province"`
	City               string `json:"city"`
	Dosis              string `json:"dosis"`
}

type GetCity struct {
	City string `json:"city"`
}
