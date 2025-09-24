package bioskopservice

import "bioskop/dto"

type BioskopService interface {
	Create(payload *dto.CreateRequest) (*dto.BioskopResponse, error)
	Get() (*dto.GetResponse, error)
	GetById(id uint) (*dto.BioskopResponse, error)
	Update(id uint, payload *dto.UpdateRequest) (*dto.BioskopResponse, error)
	Delete(id uint) error
}
