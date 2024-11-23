package usecase

import (
	"context"
	"neuro-most/tags-service/internal/entities"
)

type (
	UpdateTagUseCase interface {
		Execute(ctx context.Context, input UpdateTagInput) error
	}

	UpdateTagInput struct {
		Id   int64
		Name *string
	}

	updateTagInteractor struct {
		repo entities.TagsRepo
	}
)

func NewUpdateTagInteractor(repo entities.TagsRepo) UpdateTagUseCase {
	return &updateTagInteractor{repo: repo}
}

func (uc updateTagInteractor) Execute(ctx context.Context, input UpdateTagInput) error {
	tag, err := uc.repo.GetByID(ctx, input.Id)
	if err != nil {
		return err
	}

	if input.Name != nil {
		tag.SetName(*input.Name)
	}

	if err := uc.repo.Update(ctx, tag); err != nil {
		return err
	}

	return nil
}
