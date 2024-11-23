package usecase

import (
	"context"
	"neuro-most/tags-service/internal/entities"
)

type (
	GetByIdTagUseCase interface {
		Execute(ctx context.Context, input GetByIdTagInput) (GetByIdTagOutput, error)
	}

	GetByIdTagInput struct {
		Id int64
	}

	GetByIdTagOutput struct {
		Id   int64
		Name string
	}

	GetByIdTagPresenter interface {
		Output(tag entities.Tags) GetByIdTagOutput
	}

	getByIdTagInteractor struct {
		repo      entities.TagsRepo
		presenter GetByIdTagPresenter
	}
)

func NewGetByIdTagInteractor(repo entities.TagsRepo, presenter GetByIdTagPresenter) GetByIdTagUseCase {
	return &getByIdTagInteractor{repo: repo, presenter: presenter}
}

func (uc getByIdTagInteractor) Execute(ctx context.Context, input GetByIdTagInput) (GetByIdTagOutput, error) {
	tag, err := uc.repo.GetByID(ctx, input.Id)
	if err != nil {
		return GetByIdTagOutput{}, err
	}
	return uc.presenter.Output(tag), nil
}
