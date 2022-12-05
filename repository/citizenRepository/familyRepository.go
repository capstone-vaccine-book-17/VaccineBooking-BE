package citizenRepository

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/model"
	"time"
)

// TODO Create Family Members
func (u *citizenRepository) CreateFamilyMember(payloads citizenDto.FamilyReq) error {
	relation := model.FamilyAs{
		Name:      payloads.Relation,
		CreatedAT: time.Now(),
	}

	err := u.db.Create(&relation).Error
	if err != nil {
		return err
	}

	family := model.FamilyMember{
		CitizenId:  payloads.CitizenId,
		Name:       payloads.Name,
		Nik:        payloads.Nik,
		FamilyAsId: relation.FamilyAsID,
		Age:        payloads.Age,
		Gender:     payloads.Gender,
		CreatedAT:  time.Now(),
		UpdatedAT:  time.Now(),
	}

	errB := u.db.Create(&family).Error
	if errB != nil {
		return errB
	}

	return nil

}

// TODO GET Family Members
func (u *citizenRepository) GetFamilys(payloads citizenDto.FamilyReq) ([]citizenDto.FamilylDTO, error) {

	var members  []citizenDto.FamilylDTO

	if err := u.db.Model(&model.FamilyMember{}).Select("family_members.*").Where("citizen_id = ?", payloads.CitizenId).Find(&members).Error; err != nil {
		return members, err
	}

	return members, nil

}

// TODO Delete Member
func (u *citizenRepository) DeleteMember(payloads citizenDto.FamilylDTO) error {


	if err := u.db.Where("family_id = ?", payloads.FamilyId).Delete(&model.FamilyMember{}).Error; err != nil {
		return err
	}

	return nil
}
