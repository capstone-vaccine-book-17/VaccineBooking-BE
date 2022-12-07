package citizenDto

type BookingDto struct {
	CitizenID uint `json:"citizen_id"`
	SessionID uint `json:"session_id" validate:"required"`
	Queue     int  `json:"queue"`
}

type MaxQueue struct {
	TotalQ int `json:"total_q"`
}

type TicketBooking struct {
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
}
