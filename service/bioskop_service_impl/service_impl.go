package bioskopserviceimpl

import (
	"bioskop/dto"
	"bioskop/entity"
	bioskoprepository "bioskop/repository/bioskop_repository"
	bioskopservice "bioskop/service/bioskop_service"
	"bioskop/shared/validation"
	"errors"

	"github.com/go-playground/validator/v10"
)

type BioskopSvc struct {
	repo     bioskoprepository.BioskopRepository
	validate *validator.Validate
}

func NewBioskopService(repo bioskoprepository.BioskopRepository) bioskopservice.BioskopService {
	return &BioskopSvc{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *BioskopSvc) Create(payload *dto.CreateRequest) (*dto.BioskopResponse, error) {
	if err := s.validate.Struct(payload); err != nil {
		return nil, validation.FormatValidationError(err)
	}

	_, err := s.repo.GetByName(payload.Name)
	if err == nil {
		return nil, errors.New("film sudah ada")
	}

	request := &entity.Bioskop{
		Name:   payload.Name,
		Lokasi: payload.Lokasi,
		Rating: payload.Rating,
	}

	bioskop, errBioskop := s.repo.Create(request)
	if errBioskop != nil {
		return nil, errors.New(errBioskop.Error())
	}

	response := &dto.BioskopResponse{
		ID:     bioskop.ID,
		Name:   bioskop.Name,
		Lokasi: bioskop.Lokasi,
		Rating: bioskop.Rating,
	}

	return response, nil
}

func (s *BioskopSvc) Get() (*dto.GetResponse, error) {
	bioskop, err := s.repo.Get()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if len(bioskop) == 0 {
		return nil, errors.New("data bioskop tidak ditemukan")
	}

	response := make([]*dto.BioskopResponse, len(bioskop))
	for i, value := range bioskop {
		response[i] = &dto.BioskopResponse{
			ID:     value.ID,
			Name:   value.Name,
			Lokasi: value.Lokasi,
			Rating: value.Rating,
		}
	}

	message := &dto.GetResponse{
		Data: response,
	}

	return message, nil
}

func (s *BioskopSvc) GetById(id uint) (*dto.BioskopResponse, error) {
	bioskop, err := s.repo.GetById(id)
	if err != nil {
		return nil, errors.New("data bioskop tidak ditemukan")
	}

	response := &dto.BioskopResponse{
		ID:     bioskop.ID,
		Name:   bioskop.Name,
		Lokasi: bioskop.Lokasi,
		Rating: bioskop.Rating,
	}

	return response, nil
}

func (s *BioskopSvc) Update(id uint, payload *dto.UpdateRequest) (*dto.BioskopResponse, error) {
	if err := s.validate.Struct(payload); err != nil {
		return nil, validation.FormatValidationError(err)
	}

	_, err := s.repo.GetById(id)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	updateData := &entity.Bioskop{
		Name:   payload.Name,
		Lokasi: payload.Lokasi,
		Rating: payload.Rating,
	}

	updated := s.repo.Update(id, updateData)
	if updated != nil {
		return nil, errors.New(updated.Error())
	}

	response := &dto.BioskopResponse{
		ID:     id,
		Name:   updateData.Name,
		Lokasi: updateData.Lokasi,
		Rating: updateData.Rating,
	}

	return response, nil
}

func (s *BioskopSvc) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
