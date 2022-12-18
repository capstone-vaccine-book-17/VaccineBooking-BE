package citizenDto

type SessionDto struct {
	SessionID uint   `json:"session_id"`
	Vaccine   string `json:"vaccine"`
	Name      string `json:"name"`
	Kuota     int    `json:"kuota"`
	Dosis     string `json:"dosis"`
	Date      string `json:"date"`
	ConvDate  string `json:"conv_date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type SessionWithVaccineId struct {
	SessionID uint   `json:"session_id"`
	VaccineID uint   `json:"vaccine_id"`
	Name      string `json:"name"`
	Kuota     int    `json:"kuota"`
	Dosis     string `json:"dosis"`
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
