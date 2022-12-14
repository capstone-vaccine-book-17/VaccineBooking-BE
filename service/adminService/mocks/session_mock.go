package mocks

import "capstone_vaccine/dto/adminDto"

// Mock Create Session
func (u *MockAdmin) CreateSession(payloads adminDto.SessionRequest) (adminDto.SessionDTO, error) {
	args := u.Called(payloads)

	return args.Get(0).(adminDto.SessionDTO), args.Error(1)
}

// Mock Get All Session
func (u *MockAdmin) GetAllSession(medicalID uint) ([]adminDto.SessionWithStatusDTO, error) {
	args := u.Called(medicalID)

	return args.Get(0).([]adminDto.SessionWithStatusDTO), args.Error(1)
}

// MOCK Get Session By Id
func (u *MockAdmin) GetSessionById(payloads adminDto.SessionWithStatusDTO) (adminDto.SessionWithStatusDTO, error) {
	args := u.Called(payloads.SessionId)

	return args.Get(0).(adminDto.SessionWithStatusDTO), args.Error(1)
}

// MOCK Update Session
func (u *MockAdmin) UpdateSession(payloads adminDto.SessionRequestUpdate) (adminDto.SessionRequestUpdate, error) {
	args := u.Called(payloads)

	return args.Get(0).(adminDto.SessionRequestUpdate), args.Error(1)
}

// MOCK Delete Session
func (u *MockAdmin) DeleteSession(payloads adminDto.SessionWithStatusDTO) error {
	args := u.Called(payloads.SessionId)

	return args.Error(0)
}
