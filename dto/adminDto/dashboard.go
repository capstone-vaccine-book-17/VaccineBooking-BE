package adminDto

type CountDashboard struct {
	VaccineAvailable   string `json:"vaccine_available"`
	BookingToday       string `json:"booking_today"`
	BookingsRegistered string `json:"bookings_registered"`
}
