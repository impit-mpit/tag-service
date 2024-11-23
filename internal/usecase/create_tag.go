package usecase

import (
	"context"
	"neuro-most/tags-service/internal/entities"
)

type (
	CreateTagUseCase interface {
		Execute(ctx context.Context, input CreateTagInput) error
	}

	CreateTagInput struct {
		Name string
	}

	createTagInteractor struct {
		repo entities.TagsRepo
	}
)

func NewCreateTagInteractor(repo entities.TagsRepo) CreateTagUseCase {
	return &createTagInteractor{repo: repo}
}

func (uc createTagInteractor) Execute(ctx context.Context, input CreateTagInput) error {
	tag := entities.NewTagCreate(input.Name)
	if err := uc.repo.Create(ctx, tag); err != nil {
		return err
	}
	return nil
}
