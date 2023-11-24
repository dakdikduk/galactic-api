package usecase

import (
	"context"

	"github.com/dakdikduk/galactic-api/domain"
)

type SpacecraftUseCase struct {
	spacecraftRepo domain.SpacecraftRepository
}

func NewSpacecraftUseCase(spacecraftRepo domain.SpacecraftRepository) domain.SpacecraftUseCase {
	return &SpacecraftUseCase{
		spacecraftRepo: spacecraftRepo,
	}
}

func (uc *SpacecraftUseCase) List(ctx context.Context, params domain.ListSpacecraftParams) (res []domain.Spacecraft, err error) {
	return uc.spacecraftRepo.List(ctx, params)
}

func (uc *SpacecraftUseCase) GetByID(ctx context.Context, id int64) (domain.Spacecraft, error) {
	return uc.spacecraftRepo.GetByID(ctx, id)
}

func (uc *SpacecraftUseCase) Create(ctx context.Context, spacecraft domain.Spacecraft) {
	uc.spacecraftRepo.Create(ctx, spacecraft)
}

func (uc *SpacecraftUseCase) Update(ctx context.Context, spacecraft domain.Spacecraft) {
	uc.spacecraftRepo.Update(ctx, spacecraft)
}

func (uc *SpacecraftUseCase) Delete(ctx context.Context, id int64) {
	uc.spacecraftRepo.Delete(ctx, id)
}
