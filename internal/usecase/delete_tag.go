package usecase

import (
	"context"
	"neuro-most/tags-service/internal/entities"
)

type (
	DeleteTagUseCase interface {
		Execute(ctx context.Context, input DeleteTagInput) error
	}

	DeleteTagInput struct {
		Id int64
	}

	deleteTagInteractor struct {
		repo entities.TagsRepo
	}
)

func NewDeleteTagInteractor(repo entities.TagsRepo) DeleteTagUseCase {
	return &deleteTagInteractor{repo: repo}
}

func (uc deleteTagInteractor) Execute(ctx context.Context, input DeleteTagInput) error {
	tag, err := uc.repo.GetByID(ctx, input.Id)
	if err != nil {
		return err
	}
	if err := uc.repo.Delete(ctx, tag); err != nil {
		return err
	}
	return nil
}
