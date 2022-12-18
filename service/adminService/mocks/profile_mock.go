package mocks

import (
	"capstone_vaccine/dto/adminDto"
	"mime/multipart"
)

//M
func (u *MockAdmin) GetProfile(payloads adminDto.ProfileRequest) ([]adminDto.ProfilDTO, error) {

	args := u.Called(payloads)

	return args.Get(0).([]adminDto.ProfilDTO), args.Error(1)
}

func (u *MockAdmin) UpdateProfile(payloads adminDto.ProfileRequest) (adminDto.ProfileRequest, error) {

	args := u.Called(payloads)

	return args.Get(0).(adminDto.ProfileRequest), args.Error(1)
}

func (u *MockAdmin) UpdateImage(payloads adminDto.ProfileRequest, file multipart.File) error {

	args := u.Called(payloads)

	return args.Error(0)
}
