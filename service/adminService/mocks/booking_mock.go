package mocks

import "capstone_vaccine/dto/adminDto"

// Mock Create Booking
func (u *MockAdmin) CreateBooking(payloads adminDto.BookingDto) (adminDto.BookingDto, error) {
	args := u.Called(payloads)

	return args.Get(0).(adminDto.BookingDto), args.Error(1)
}

// Mock Update Booking
func (u *MockAdmin) UpdateBooking(payloads adminDto.UpdateBooking) (adminDto.UpdateBooking, error) {
	args := u.Called(payloads)

	return args.Get(0).(adminDto.UpdateBooking), args.Error(1)
}

// MOCK Get All Booking
func (u *MockAdmin) GetAllBooking(medicalId uint) ([]adminDto.BookingAllDto, error) {
	args := u.Called(medicalId)

	return args.Get(0).([]adminDto.BookingAllDto), args.Error(1)
}

// MOCK Get By Id Booking
func (u *MockAdmin) GetBookingById(payloads adminDto.BookingAllDto) (adminDto.BookingAllDto, error) {
	args := u.Called(payloads.BookingId)

	return args.Get(0).(adminDto.BookingAllDto), args.Error(1)
}

// MOCK Delete Booking
func (u *MockAdmin) DeleteBooking(payloads adminDto.BookingAllDto) error {
	args := u.Called(payloads.BookingId)

	return args.Error(0)
}
