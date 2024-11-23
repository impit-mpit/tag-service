package presenter

import (
	"neuro-most/tags-service/internal/entities"
	"neuro-most/tags-service/internal/usecase"
)

type GetByIdPresenter struct {
}

func NewGetByIdPresenter() GetByIdPresenter {
	return GetByIdPresenter{}
}

func (p GetByIdPresenter) Output(tag entities.Tags) usecase.GetByIdTagOutput {
	return usecase.GetByIdTagOutput{
		Id:   tag.ID(),
		Name: tag.Name(),
	}
}
