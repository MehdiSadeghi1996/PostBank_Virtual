package repository

import (
	"PostBank_Virtual_Banking/domain"
	"gorm.io/gorm"
	"time"
)

type ProceedingRepository struct {
	db *gorm.DB
}

func NewProceedingRepository(db *gorm.DB) *ProceedingRepository {
	return &ProceedingRepository{db: db}
}

func (pr *ProceedingRepository) CreateProceeding(proceeding *domain.Proceeding) error {
	result := pr.db.Create(proceeding)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *ProceedingRepository) GetProceedingByID(id uint) (*domain.Proceeding, error) {
	var user domain.Proceeding
	result := pr.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (pr *ProceedingRepository) GetProceedingWithPagination(page int, pageSize int) ([]domain.Proceeding, error) {
	var proceedings []domain.Proceeding
	offset := page

	result := pr.db.Offset(offset).Limit(pageSize).Find(&proceedings)
	if result.Error != nil {
		return nil, result.Error
	}

	return proceedings, nil
}

func (pr *ProceedingRepository) MultipleColumnFilter(filters map[string]interface{}, startTime time.Time, endTime time.Time) ([]domain.Proceeding, error) {
	var proceedings []domain.Proceeding

	query := pr.db
	for column, value := range filters {
		query = query.Where(column+" = ?", value)
	}
	query = query.Where("created_at >= ? AND created_at <= ?", startTime, endTime)

	result := query.Find(&proceedings)
	if result.Error != nil {
		return nil, result.Error
	}

	return proceedings, nil
}

func (pr *ProceedingRepository) UpdateProceeding(proceeding *domain.Proceeding) error {
	result := pr.db.Save(proceeding)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *ProceedingRepository) DeleteProceeding(proceeding *domain.Proceeding) error {
	result := pr.db.Delete(proceeding)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
