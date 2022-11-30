package adminDto

type MedicalDto struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Province string `json:"province" form:"province"`
	PostCode string `json:"post_code" form:"post_code"`
	Country  string `json:"country" form:"country"`
	City     string `json:"city" form:"city"`
	NoTlp    string `json:"no_tlp" form:"no_tlp"`
}
