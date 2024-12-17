package service

import (
	"moydom_api/internal/domain"
)

type PropertyService struct {
	propertyRepo domain.PropertyRepository
}

func NewPropertyService(propertyRepo domain.PropertyRepository) *PropertyService {
	return &PropertyService{propertyRepo: propertyRepo}
}

func (s *PropertyService) GetAllProperties(page, pageSize int) ([]domain.Property, int64, error) {
	return s.propertyRepo.GetAll(page, pageSize)
}

//func (s *PropertyService) CreateProperty(input domain.Property) error {
//	propertyExist, err := s.propertyRepo.FindDuplicate(input)
//	if err != nil || propertyExist {
//		return errors.New("Недвижимость существует")
//	}
//	return nil
//}
