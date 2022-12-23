package citizenService

import "capstone_vaccine/dto/citizenDto"

//TODO Create Family Member

func (s *citizenService) CreateFamilyMember(payloads citizenDto.FamilyReq) error {
	return s.citizenRepository.CreateFamilyMember(payloads)

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
			FamilyAs: v.FamilyAs,
			Gender:   v.Gender,
		})
	}

	return members, nil

}

//TODO Delete Members
func (s *citizenService) DeleteMember(payloads citizenDto.FamilylDTO) error {
	return s.citizenRepository.DeleteMember(payloads)
}

//TODO GET Detail Member
func (s *citizenService) GetDetailMember(payload citizenDto.FamilylDTO) (citizenDto.FamilylDTO, error) {
	return s.citizenRepository.GetDetailMember(payload)

}
