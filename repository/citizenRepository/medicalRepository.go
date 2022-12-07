package citizenRepository

import (
	"capstone_vaccine/dto/citizenDto"
	"capstone_vaccine/model"
)

// TODO GET CITY CITIZEN
func (u *citizenRepository) GetCityCitizen(id uint) (citizenDto.GetCity, error) {
	var citizen citizenDto.GetCity

	if err := u.db.Model(&model.Citizen{}).Select("addresses.city").Joins("join addresses on addresses.address_id = citizens.address_id").Where("citizens.citizen_id = ?", id).Find(&citizen).Error; err != nil {
		return citizen, err
	}

	return citizen, nil
}

// TODO GET ALL MEDICAL BY CITY AND BY SEARCH
func (u *citizenRepository) GetMedicalByCity(payloads citizenDto.SearchKey) ([]citizenDto.SearchDto, error) {
	var search []citizenDto.SearchDto

	if payloads.S != "" && payloads.Q != "" {
		if err := u.db.Raw("SELECT medical_facilitys.medical_facilitys_id, medical_facilitys.image, medical_facilitys.name, addresses.address, addresses.province, addresses.city, GROUP_CONCAT(DISTINCT sessions.dosis ORDER BY sessions.session_id ASC) as dosis FROM medical_facilitys JOIN addresses on addresses.address_id = medical_facilitys.address_id JOIN sessions ON sessions.medical_facilitys_id = medical_facilitys.medical_facilitys_id WHERE sessions.status = 'process' AND medical_facilitys.name LIKE ? OR addresses.city LIKE ? AND sessions.dosis = ? GROUP BY medical_facilitys.name", "%"+payloads.S+"%", "%"+payloads.S+"%", payloads.Q).Scan(&search).Error; err != nil {
			return nil, err
		}
	} else if payloads.S != "" {
		if err := u.db.Raw("SELECT medical_facilitys.medical_facilitys_id, medical_facilitys.image, medical_facilitys.name, addresses.address, addresses.province, addresses.city, GROUP_CONCAT(DISTINCT sessions.dosis ORDER BY sessions.session_id ASC) as dosis FROM medical_facilitys JOIN addresses on addresses.address_id = medical_facilitys.address_id JOIN sessions ON sessions.medical_facilitys_id = medical_facilitys.medical_facilitys_id WHERE sessions.status = 'process' AND medical_facilitys.name LIKE ? OR addresses.city LIKE ? GROUP BY medical_facilitys.name", "%"+payloads.S+"%", "%"+payloads.S+"%").Scan(&search).Error; err != nil {
			return nil, err
		}
	} else if payloads.Q != "" {
		if err := u.db.Raw("SELECT medical_facilitys.medical_facilitys_id, medical_facilitys.image, medical_facilitys.name, addresses.address, addresses.province, addresses.city, GROUP_CONCAT(DISTINCT sessions.dosis ORDER BY sessions.session_id ASC) as dosis FROM medical_facilitys JOIN addresses on addresses.address_id = medical_facilitys.address_id JOIN sessions ON sessions.medical_facilitys_id = medical_facilitys.medical_facilitys_id WHERE sessions.status = 'process' AND addresses.city = ? AND sessions.dosis = ? GROUP BY medical_facilitys.name", payloads.City, payloads.Q).Scan(&search).Error; err != nil {
			return nil, err
		}
	} else {
		if payloads.City != "" {
			if err := u.db.Raw("SELECT medical_facilitys.medical_facilitys_id, medical_facilitys.image, medical_facilitys.name, addresses.address, addresses.province, addresses.city, GROUP_CONCAT(DISTINCT sessions.dosis ORDER BY sessions.session_id ASC) as dosis FROM medical_facilitys JOIN addresses on addresses.address_id = medical_facilitys.address_id JOIN sessions ON sessions.medical_facilitys_id = medical_facilitys.medical_facilitys_id WHERE sessions.status = 'process' AND addresses.city = ? GROUP BY medical_facilitys.name", payloads.City).Scan(&search).Error; err != nil {
				return nil, err
			}
		} else {
			if err := u.db.Raw("SELECT medical_facilitys.medical_facilitys_id, medical_facilitys.image, medical_facilitys.name, addresses.address, addresses.province, addresses.city, GROUP_CONCAT(DISTINCT sessions.dosis ORDER BY sessions.session_id ASC) as dosis FROM medical_facilitys JOIN addresses on addresses.address_id = medical_facilitys.address_id JOIN sessions ON sessions.medical_facilitys_id = medical_facilitys.medical_facilitys_id WHERE sessions.status = 'process' GROUP BY medical_facilitys.name").Scan(&search).Error; err != nil {
				return nil, err
			}
		}
	}

	return search, nil
}

// TODO GET MEDICAL BY ID
func (u *citizenRepository) GetMedicalById(medicalID uint) (citizenDto.MedicalDto, error) {
	var medical citizenDto.MedicalDto
	if err := u.db.Model(&model.MedicalFacilitys{}).Select("medical_facilitys.*, addresses.*").
		Joins("JOIN addresses ON addresses.address_id = medical_facilitys.address_id").
		Where("medical_facilitys.medical_facilitys_id = ?", medicalID).Find(&medical).Error; err != nil {
		return medical, err
	}

	return medical, nil
}
