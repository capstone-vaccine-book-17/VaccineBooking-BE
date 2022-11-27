package adminDto

type SessionDTO struct {
	Name      string `json:"name" validate:"required"`
	VaccineId uint   `json:"vaccine_id"`
	StartTime string `json:"startTime" validate:"required"`
	Kuota     int    `json:"kuota" validate:"required"`
	Dosis     string `json:"dosis" validate:"required"`
	EndTime   string `json:"endTime" validate:"required"`
	Date      string `json:"date" validate:"required"`
}

type SessionWithStatusDTO struct {
	SessionId   uint   `json:"session_id"`
	Name        string `json:"name" validate:"required"`
	VaccineName string `json:"vaccine_name"`
	StartTime   string `json:"startTime" validate:"required"`
	Kuota       int    `json:"kuota" validate:"required"`
	Dosis       string `json:"dosis" validate:"required"`
	EndTime     string `json:"endTime" validate:"required"`
	Date        string `json:"date" validate:"required"`
	Status      string `json:"status"`
}

type SessionRequest struct {
	Name               string `json:"name" validate:"required"`
	MedicalFacilitysId uint   `json:"medical_facilitys_id"`
	VaccineId          uint   `json:"vaccine_id"`
	StartTime          string `json:"startTime" validate:"required"`
	Kuota              int    `json:"kuota" validate:"required"`
	Dosis              string `json:"dosis" validate:"required"`
	EndTime            string `json:"endTime" validate:"required"`
	Date               string `json:"date" validate:"required"`
}

type SessionRequestUpdate struct {
	SessionId uint   `json:"session_id"`
	Name      string `json:"name" validate:"required"`
	VaccineId uint   `json:"vaccine_id"`
	StartTime string `json:"startTime" validate:"required"`
	Kuota     int    `json:"kuota" validate:"required"`
	Dosis     string `json:"dosis" validate:"required"`
	EndTime   string `json:"endTime" validate:"required"`
}

type CountKuota struct {
	TotalS int `json:"total_s"`
	TotalV int `json:"total_v"`
}
