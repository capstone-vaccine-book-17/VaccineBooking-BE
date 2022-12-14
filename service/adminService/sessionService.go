package adminService

import (
	"capstone_vaccine/dto/adminDto"
	"errors"
	"time"
)

// TODO CREATE SESSION
func (s *adminService) CreateSession(payloads adminDto.SessionRequest) (adminDto.SessionDTO, error) {
	var dto adminDto.SessionDTO

	kuota, _ := s.adminRepository.CountKuota(payloads.VaccineId)
	if kuota.TotalS >= kuota.TotalV {
		return dto, errors.New("kuota vaksin yang di input melebihi batas")
	} else if payloads.Kuota+kuota.TotalS > kuota.TotalV {
		return dto, errors.New("kuota vaksin yang di input melebihi batas")
	}

	temp := adminDto.SessionRequest{
		Name:               payloads.Name,
		MedicalFacilitysId: payloads.MedicalFacilitysId,
		VaccineId:          payloads.VaccineId,
		StartTime:          payloads.StartTime,
		Kuota:              payloads.Kuota,
		Dosis:              payloads.Dosis,
		EndTime:            payloads.EndTime,
		Date:               payloads.Date,
	}

	res, err := s.adminRepository.CreateSession(temp)
	if err != nil {
		return res, err
	}
	return res, nil
}

func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

// TODO GET ALL SESSION
func (s *adminService) GetAllSession() ([]adminDto.SessionWithStatusDTO, error) {

	// Set time today with date format
	today := time.Now()
	dateFormat := today.Format("2006-01-02")
	convDate := string(dateFormat)
	loc := today.Location()

	res, err := s.adminRepository.GetAllSession()
	if err != nil {
		return nil, err
	}

	// loop data from session and check if date and time same with convDate and different time or not
	for i := range res {
		pasttime, _ := time.ParseInLocation("15:04", res[i].EndTime, loc)
		diff := today.Sub(pasttime)
		if res[i].Date <= convDate {
			if diff.Hours() != 0 {
				err := s.adminRepository.AutoUpdateSession(res[i].Date, res[i].EndTime)
				if err != nil {
					return res, err
				}
			}
		}

	}

	return res, nil
}

// TODO GET SESSION BY ID
func (s *adminService) GetSessionById(payloads adminDto.SessionWithStatusDTO) (adminDto.SessionWithStatusDTO, error) {
	res, err := s.adminRepository.GetSessionById(payloads)

	if res.SessionId < 1 {
		return res, errors.New("record not found")
	}

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *adminService) UpdateSession(payloads adminDto.SessionRequestUpdate) (adminDto.SessionRequestUpdate, error) {
	// DTO for get data kuota on session table by id
	dto_sessionById := adminDto.SessionWithStatusDTO{
		SessionId: payloads.SessionId,
	}

	dto_update := adminDto.SessionRequestUpdate{}

	kuotaById, _ := s.adminRepository.GetSessionById(dto_sessionById)

	kuota, _ := s.adminRepository.CountKuota(payloads.VaccineId)
	// Check if kuota already maximum
	kuota.TotalS = kuota.TotalS - kuotaById.Kuota

	if kuota.TotalS >= kuota.TotalV {
		return dto_update, errors.New("kuota vaksin yang di input melebihi batas")
	} else if payloads.Kuota+kuota.TotalS > kuota.TotalV {
		return dto_update, errors.New("kuota vaksin yang di input melebihi batas")
	}

	res, err := s.adminRepository.UpdateSession(payloads)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *adminService) DeleteSession(payloads adminDto.SessionWithStatusDTO) error {

	err := s.adminRepository.DeleteSession(payloads)

	if err != nil {
		return err
	}

	return nil
}
