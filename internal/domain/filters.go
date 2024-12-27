package domain

type FilterDistrict struct {
	ID   uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name"`
}

type FilterMicrodistrict struct {
	ID   uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name"`
}

type FilterRegion struct {
	ID   uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name"`
}

type FiltersRepository interface {
	GetDistricts() ([]FilterDistrict, error)
	AddDistrict(input FilterDistrict) (FilterDistrict, error)
	UpdateDistrict(id int, newName string) (FilterDistrict, error)
	DeleteDistrict(id int) error

	GetMicrodistricts() ([]FilterMicrodistrict, error)
	AddMicrodistrict(input FilterMicrodistrict) (FilterMicrodistrict, error)
	UpdateMicrodistrict(id int, newName string) (FilterMicrodistrict, error)
	DeleteMicrodistrict(id int) error

	GetRegions() ([]FilterRegion, error)
	AddRegion(input FilterRegion) (FilterRegion, error)
	UpdateRegion(id int, newName string) (FilterRegion, error)
	DeleteRegion(id int) error
}
