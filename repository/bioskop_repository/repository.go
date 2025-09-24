package bioskoprepository

import "bioskop/entity"

type BioskopRepository interface {
	Create(bioskop *entity.Bioskop) (*entity.Bioskop, error)
	Get() ([]entity.Bioskop, error)
	GetByName(name string) (*entity.Bioskop, error)
	GetById(id uint) (*entity.Bioskop, error)
	Update(id uint, bioskop *entity.Bioskop) error
	Delete(id uint) error
}
