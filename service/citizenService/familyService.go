package citizenService

import "capstone_vaccine/dto/citizenDto"

//TODO Create Family Member

func (s *citizenService) CreateFamilyMember(payloads citizenDto.FamilyReq) error {
	err := s.citizenRepository.CreateFamilyMember(payloads)

	if err != nil {
		return err
	}

	return nil

}
