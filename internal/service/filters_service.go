package service

import "moydom_api/internal/domain"

type FiltersService struct {
	filtersRepo domain.FiltersRepository
}

func NewFiltersService(filtersRepo domain.FiltersRepository) *FiltersService {
	return &FiltersService{filtersRepo: filtersRepo}
}

// Districts

func (s *FiltersService) GetDistricts() ([]domain.FilterDistrict, error) {
	return s.filtersRepo.GetDistricts()
}

func (s *FiltersService) AddDistrict(input domain.FilterDistrict) (domain.FilterDistrict, error) {
	// Можно как-то провалидировать
	return s.filtersRepo.AddDistrict(input)
}

func (s *FiltersService) UpdateDistrict(id int, newName string) (domain.FilterDistrict, error) {
	return s.filtersRepo.UpdateDistrict(id, newName)
}

func (s *FiltersService) DeleteDistrict(id int) error {
	return s.filtersRepo.DeleteDistrict(id)
}

// MicroDistricts

func (s *FiltersService) GetMicrodistricts() ([]domain.FilterMicrodistrict, error) {
	return s.filtersRepo.GetMicrodistricts()
}

func (s *FiltersService) AddMicrodistrict(input domain.FilterMicrodistrict) (domain.FilterMicrodistrict, error) {
	// Можно как-то провалидировать
	return s.filtersRepo.AddMicrodistrict(input)
}

func (s *FiltersService) UpdateMicrodistrict(id int, newName string) (domain.FilterMicrodistrict, error) {
	return s.filtersRepo.UpdateMicrodistrict(id, newName)
}

func (s *FiltersService) DeleteMicrodistrict(id int) error {
	return s.filtersRepo.DeleteMicrodistrict(id)
}
