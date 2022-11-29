package adminDto

type BookingDto struct {
	CitizenId uint   `json:"citizen_id"`
	SessionId uint   `json:"session_id"`
	Nama      string `json:"nama" validate:"required"`
	Nik       string `json:"nik" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Queue     int    `json:"queue"`
}

type UpdateSessionBooking struct {
	SessionId uint `json:"session_id"`
	Kuota     int  `json:"kuota"`
}

type MaxQueue struct {
	TotalQ int `json:"total_q"`
}
