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
