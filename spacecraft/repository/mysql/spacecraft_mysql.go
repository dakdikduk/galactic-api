package repository

import (
	"context"

	"github.com/dakdikduk/galactic-api/domain"
	"github.com/jinzhu/gorm"
)

type MySQLSpacecraftRepository struct {
	DB *gorm.DB
}

func NewMySQLSpacecraftRepository(db *gorm.DB) domain.SpacecraftRepository {
	return &MySQLSpacecraftRepository{
		DB: db,
	}
}

func (r *MySQLSpacecraftRepository) List(ctx context.Context, params domain.ListSpacecraftParams) (res []domain.Spacecraft, err error) {
	query := r.DB
	if params.Name != "" {
		query = query.Where("name = ?", params.Name)
	}
	if params.Class != "" {
		query = query.Where("class = ?", params.Class)
	}
	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}
	if params.Page > 0 && params.PerPage > 0 {
		offset := (params.Page - 1) * params.PerPage
		query = query.Offset(offset).Limit(params.PerPage)
	}

	err = query.Find(&res).Error
	return res, err
}

func (r *MySQLSpacecraftRepository) GetByID(ctx context.Context, id int64) (domain.Spacecraft, error) {
	var spacecraft domain.Spacecraft
	err := r.DB.Where("id = ?", id).First(&spacecraft).Error
	return spacecraft, err
}

func (r *MySQLSpacecraftRepository) Create(ctx context.Context, spacecraft domain.Spacecraft) {
	r.DB.Create(&spacecraft)
}

func (r *MySQLSpacecraftRepository) Update(ctx context.Context, spacecraft domain.Spacecraft) {
	r.DB.Save(&spacecraft)
}

func (r *MySQLSpacecraftRepository) Delete(ctx context.Context, id int64) {
	r.DB.Where("id = ?", id).Delete(&domain.Spacecraft{})
}
