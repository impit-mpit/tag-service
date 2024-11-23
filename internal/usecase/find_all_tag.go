package usecase

import (
	"context"
	"neuro-most/tags-service/internal/entities"
)

type (
	FindAllTag interface {
		Execute(ctx context.Context, input FindAllTagInput) ([]FindAllTagOutput, int64, error)
	}

	FindAllTagInput struct {
		Page     int64
		PageSize int64
	}

	FindAllTagOutput struct {
		Id   int64
		Name string
	}

	FindAllTagPresenter interface {
		Output(tags []entities.Tags) []FindAllTagOutput
	}

	findAllTagInteractor struct {
		repo      entities.TagsRepo
		presenter FindAllTagPresenter
	}
)

func NewFindAllTagInteractor(repo entities.TagsRepo, presenter FindAllTagPresenter) FindAllTag {
	return &findAllTagInteractor{repo: repo, presenter: presenter}
}

func (uc findAllTagInteractor) Execute(ctx context.Context, input FindAllTagInput) ([]FindAllTagOutput, int64, error) {
	tags, total, err := uc.repo.Fetch(ctx, input.Page, input.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return uc.presenter.Output(tags), total, nil
}
