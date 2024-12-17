package repository

import (
	"gorm.io/gorm"
	"moydom_api/internal/domain"
)

type PropertyRepository struct {
	db *gorm.DB
}

func NewPropertyRepository(db *gorm.DB) *PropertyRepository {
	return &PropertyRepository{db: db}
}

func (r *PropertyRepository) GetAll(page, limit int) ([]domain.Property, int64, error) {
	var properties []domain.Property
	var total int64

	offset := (page - 1) * limit

	if err := r.db.Model(&domain.Property{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.Limit(limit).Offset(offset).Find(&properties).Error; err != nil {
		return nil, 0, err
	}
	return properties, total, nil
}
