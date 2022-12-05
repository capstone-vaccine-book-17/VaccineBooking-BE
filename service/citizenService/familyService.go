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

//TODO GET Family Members
func (s *citizenService) GetFamilys(payloads citizenDto.FamilyReq) ([]citizenDto.FamilylDTO, error) {
	var members []citizenDto.FamilylDTO

	res, err := s.citizenRepository.GetFamilys(payloads)

	if err != nil {
		return res, err
	}

	for _, v := range res {
		members = append(members, citizenDto.FamilylDTO{
			FamilyId: v.FamilyId,
			Name:     v.Name,
			Nik:      v.Nik,
			Age:      v.Age,
			Gender:   v.Gender,
		})
	}

	return members, nil

}

//TODO Delete Members
func (s *citizenService) DeleteMember(payloads citizenDto.FamilylDTO) error {
	err := s.citizenRepository.DeleteMember(payloads)

	if err != nil {
		return err
	}

	return nil

}
