package domain

type Property struct {
	ID                 uint   `json:"id" gorm:"primary_key"`
	OwnerID            uint   `json:"owner_id" binding:"required"`
	Type               string `json:"type" binding:"required"`
	Address            string `json:"address" binding:"required"`
	District           string `json:"district" binding:"required"`
	Microdistrict      string `json:"microdistrict" binding:"required"`
	Description        string `json:"description" binding:"required"`
	YearOfConstruction string `json:"year_of_construction" binding:"required"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

type PropertyRepository interface {
	GetAll(page, limit int) ([]Property, int64, error)
}
