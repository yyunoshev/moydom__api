package repository

import (
	"gorm.io/gorm"
	"moydom_api/internal/domain"
)

type FiltersRepository struct {
	db *gorm.DB
}

func NewFiltersRepository(db *gorm.DB) *FiltersRepository {
	return &FiltersRepository{db: db}
}

// Districts

func (r *FiltersRepository) GetDistricts() ([]domain.FilterDistrict, error) {
	var districts []domain.FilterDistrict
	if err := r.db.Find(&districts).Error; err != nil {
		return nil, err
	}
	return districts, nil
}

func (r *FiltersRepository) AddDistrict(district domain.FilterDistrict) (domain.FilterDistrict, error) {
	if err := r.db.Create(&district).Error; err != nil {
		return domain.FilterDistrict{}, err
	}
	return district, nil
}

func (r *FiltersRepository) UpdateDistrict(id int, newName string) (domain.FilterDistrict, error) {
	var district domain.FilterDistrict
	result := r.db.Model(&district).Where("id = ?", id).Update("name", newName)
	if result.Error != nil {
		return domain.FilterDistrict{}, result.Error
	}
	if result.RowsAffected == 0 {
		return domain.FilterDistrict{}, gorm.ErrRecordNotFound
	}
	return district, nil
}

func (r *FiltersRepository) DeleteDistrict(id int) error {
	result := r.db.Where("id = ?", id).Delete(&domain.FilterDistrict{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// MicroDistricts

func (r *FiltersRepository) GetMicrodistricts() ([]domain.FilterMicrodistrict, error) {
	var microdistricts []domain.FilterMicrodistrict
	if err := r.db.Find(&microdistricts).Error; err != nil {
		return nil, err
	}
	return microdistricts, nil
}

func (r *FiltersRepository) AddMicrodistrict(microdistrict domain.FilterMicrodistrict) (domain.FilterMicrodistrict, error) {
	if err := r.db.Create(&microdistrict).Error; err != nil {
		return domain.FilterMicrodistrict{}, err
	}
	return microdistrict, nil
}

func (r *FiltersRepository) UpdateMicrodistrict(id int, newName string) (domain.FilterMicrodistrict, error) {
	var microdistrict domain.FilterMicrodistrict
	result := r.db.Model(&microdistrict).Where("id = ?", id).Update("name", newName)
	if result.Error != nil {
		return domain.FilterMicrodistrict{}, result.Error
	}
	if result.RowsAffected == 0 {
		return domain.FilterMicrodistrict{}, gorm.ErrRecordNotFound
	}
	return microdistrict, nil
}

func (r *FiltersRepository) DeleteMicrodistrict(id int) error {
	result := r.db.Where("id = ?", id).Delete(&domain.FilterMicrodistrict{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
