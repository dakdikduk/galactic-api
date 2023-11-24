package domain

import (
	"context"
	"time"
)

type Spacecraft struct {
	ID        uint       `gorm:"primary_key"`
	Name      string     `gorm:"not null"`
	Class     string     `gorm:"not null"`
	Armament  []Armament `gorm:"type:JSON;not null"`
	Crew      int        `gorm:"not null"`
	Image     string     `gorm:"default:null"`
	Value     float64    `gorm:"not null"`
	Status    string     `gorm:"not null"`
	CreatedAt time.Time  `gorm:"default:current_timestamp"`
	UpdatedAt time.Time  `gorm:"default:current_timestamp"`
}

type ListSpacecraftParams struct {
	Name    string
	Class   string
	Status  string
	Page    int
	PerPage int
}

type Armament struct {
	Title string `json:"title"`
	Qty   string `json:"qty"`
}

type SpacecraftRepository interface {
	List(ctx context.Context, params ListSpacecraftParams) (res []Spacecraft, err error)
	GetByID(ctx context.Context, id int64) (Spacecraft, error)
	Create(ctx context.Context, spacecraft Spacecraft)
	Update(ctx context.Context, spacecraft Spacecraft)
	Delete(ctx context.Context, id int64)
}

type SpacecraftUseCase interface {
	List(ctx context.Context, params ListSpacecraftParams) (res []Spacecraft, err error)
	GetByID(ctx context.Context, id int64) (Spacecraft, error)
	Create(ctx context.Context, spacecraft Spacecraft)
	Update(ctx context.Context, spacecraft Spacecraft)
	Delete(ctx context.Context, id int64)
}
