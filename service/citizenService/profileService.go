package citizenService

import (
	"capstone_vaccine/dto/citizenDto"
)

// TODO Get Profile
func (s *citizenService) GetProfile(payloads citizenDto.ProfileReq) (citizenDto.ProfileDTO, error) {

	res, err := s.citizenRepository.GetProfile(payloads)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *citizenService) UploadImage(payloads citizenDto.ProfileReq) (citizenDto.ProfileReq, error) {

	res, err := s.citizenRepository.UploadImage(payloads)

	out := citizenDto.ProfileReq{
		Image: res.Image,
	}

	if err != nil {
		return out, err
	}

	return out, nil

}
