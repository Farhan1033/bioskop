package bioskoprepositorypg

import (
	"bioskop/entity"
	bioskoprepository "bioskop/repository/bioskop_repository"
	"errors"

	"gorm.io/gorm"
)

type BioskopRepo struct {
	db *gorm.DB
}

func NewBioskopRepository(db *gorm.DB) bioskoprepository.BioskopRepository {
	return &BioskopRepo{
		db: db,
	}
}

func (r *BioskopRepo) Create(bioskop *entity.Bioskop) (*entity.Bioskop, error) {
	if err := r.db.Create(bioskop).Error; err != nil {
		return nil, errors.New("internal server error")
	}

	return bioskop, nil
}
func (r *BioskopRepo) Get() ([]entity.Bioskop, error) {
	var bioskop []entity.Bioskop

	if err := r.db.Find(&bioskop).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("bioskop tidak ditemukan")
		}
		return nil, errors.New(err.Error())
	}

	return bioskop, nil
}

func (r *BioskopRepo) GetByName(name string) (*entity.Bioskop, error) {
	var bioskop entity.Bioskop

	if err := r.db.First(&bioskop, "name = ?", name).Error; err != nil {
		return nil, errors.New("film sudah ada")
	}

	return &bioskop, nil
}

func (r *BioskopRepo) GetById(id uint) (*entity.Bioskop, error) {
	var bioskop entity.Bioskop

	if err := r.db.First(&bioskop, "id = ?", id).Error; err != nil {
		return nil, errors.New("film sudah ada")
	}

	return &bioskop, nil
}

func (r *BioskopRepo) Update(id uint, bioskop *entity.Bioskop) error {
	updates := map[string]interface{}{
		"name":   bioskop.Name,
		"lokasi": bioskop.Lokasi,
		"rating": bioskop.Rating,
	}

	result := r.db.Model(&entity.Bioskop{}).Where("id = ?", id).Updates(updates)

	if result.Error != nil {
		return errors.New("gagal saat update data")
	}

	if result.RowsAffected == 0 {
		return errors.New("tidak ada data yang terupdate")
	}

	return nil
}
func (r *BioskopRepo) Delete(id uint) error {
	if err := r.db.Delete(&entity.Bioskop{}, "id = ?", id).Error; err != nil {
		return errors.New("gagal saat menghapus data")
	}

	return nil
}
