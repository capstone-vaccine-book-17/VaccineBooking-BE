package adminService

import (
	"capstone_vaccine/dto/adminDto"
	"capstone_vaccine/utils"
	"context"
	"errors"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"golang.org/x/crypto/bcrypt"
)

// TODO Get Profile Admin
func (s *adminService) GetProfile(payloads adminDto.ProfileRequest) ([]adminDto.ProfilDTO, error) {

	var profile []adminDto.ProfilDTO

	res, err := s.adminRepository.GetProfile(payloads)

	if err != nil {
		return nil, err
	}
	for _, p := range res {

		profile = append(profile, adminDto.ProfilDTO{
			Name:              p.Name,
			Image:             p.Image,
			Address:           p.Address,
			ResponsiblePerson: p.ResponsiblePerson,
			Username:          p.Username,
			Password:          p.Password,
		})
	}
	return profile, nil
}

// TODO Update Profile & Change Password
func (s *adminService) UpdateProfile(payloads adminDto.ProfileRequest) (adminDto.ProfileRequest, error) {

	if payloads.NewPassword == "" {
		payloads.NewPassword = payloads.Password

	}
	hash, _ := utils.HashBcrypt(payloads.NewPassword)

	dto := adminDto.ProfileRequest{
		AdminID:            payloads.AdminID,
		MedicalFacilitysId: payloads.MedicalFacilitysId,
		Name:               payloads.Name,
		Image:              payloads.Image,
		Address:            payloads.Address,
		ResponsiblePerson:  payloads.ResponsiblePerson,
		Username:           payloads.Username,
		Password:           payloads.Password,
		NewPassword:        hash,
	}
	new, _ := s.adminRepository.GetAdmin(payloads)

	if err := bcrypt.CompareHashAndPassword([]byte(new.Password), []byte(dto.Password)); err != nil {

		return dto, errors.New("password incorrect")

	}

	_, err := s.adminRepository.UpdateProfile(dto)

	if err != nil {
		return dto, err
	}

	return dto, nil
}

// TODO Uploud Image
func (s *adminService) UpdateImage(payloads adminDto.ProfileRequest, file multipart.File) error {

	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	ctx := context.Background()

	result, errs := cld.Upload.Upload(ctx, file, uploader.UploadParams{})
	if errs != nil {
		return errs
	}

	temp := adminDto.ProfileRequest{
		MedicalFacilitysId: payloads.MedicalFacilitysId,
		Image:              result.SecureURL,
	}

	err := s.adminRepository.UpdateImage(temp)
	if err != nil {
		return err
	}

	return nil

}
