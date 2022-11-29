package adminDto

type BookingDto struct {
	CitizenId uint   `json:"citizen_id"`
	SessionId uint   `json:"session_id"`
	Nama      string `json:"nama" validate:"required"`
	Nik       string `json:"nik" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Queue     int    `json:"queue"`
}

type BookingAllDto struct {
	BookingId   uint   `json:"booking_id"`
	CitizenName string `json:"citizen_name"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Date        string `json:"date"`
	Dosis       string `json:"dosis"`
	Nik         string `json:"nik"`
	Queue       int    `json:"queue"`
	Status      string `json:"status"`
}
type UpdateBooking struct {
	BookingId uint   `json:"booking_id"`
	Status    string `json:"status" validate:"required"`
}
type UpdateSessionBooking struct {
	SessionId uint `json:"session_id"`
	Kuota     int  `json:"kuota"`
}

type MaxQueue struct {
	TotalQ int `json:"total_q"`
}
