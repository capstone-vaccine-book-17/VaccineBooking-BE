package citizenDto

type TicketDetails struct {
	Queue     string `json:"queue"`
	Name      string `json:"name"`
	Nik       string `json:"nik"`
	Gender    string `json:"gender"`
	Vaccine   string `json:"vaccine"`
	Dosis     string `json:"dosis"`
	Date      string `json:"date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	RsName    string `json:"rs_name"`
	Status    string `json:"status"`
}
