package repository

import (
	"PostBank_Virtual_Banking/domain"
	"gorm.io/gorm"
	"time"
)

type GamaRepository struct {
	db *gorm.DB
}

func NewGamaRepository(db *gorm.DB) *GamaRepository {
	return &GamaRepository{db: db}
}

func (gr *GamaRepository) Create(gama *domain.Gama) error {
	result := gr.db.Create(gama)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *GamaRepository) GetByID(id uint) (*domain.Gama, error) {
	var gama domain.Gama
	result := gr.db.First(&gama, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &gama, nil
}

func (gr *GamaRepository) GetWithPagination(page int, pageSize int) ([]domain.Gama, error) {
	var gama []domain.Gama
	offset := page

	result := gr.db.Offset(offset).Limit(pageSize).Find(&gama)
	if result.Error != nil {
		return nil, result.Error
	}

	return gama, nil
}

func (gr *GamaRepository) MultipleColumnFilter(filters map[string]interface{}, startTime time.Time, endTime time.Time) ([]domain.Gama, error) {
	var gamas []domain.Gama

	query := gr.db
	for column, value := range filters {
		query = query.Where(column+" = ?", value)
	}

	query = query.Where("created_at >= ? AND created_at <= ?", startTime, endTime)

	result := query.Find(&gamas)
	if result.Error != nil {
		return nil, result.Error
	}

	return gamas, nil
}

func (gr *GamaRepository) Update(gama *domain.Gama) error {
	result := gr.db.Save(gama)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *GamaRepository) Delete(gama *domain.Gama) error {
	result := gr.db.Delete(gama)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
