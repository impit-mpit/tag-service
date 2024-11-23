package presenter

import (
	"neuro-most/tags-service/internal/entities"
	"neuro-most/tags-service/internal/usecase"
)

type FindAllTagPresenter struct {
}

func NewFindAllTagPresenter() FindAllTagPresenter {
	return FindAllTagPresenter{}
}

func (p FindAllTagPresenter) Output(tags []entities.Tags) []usecase.FindAllTagOutput {
	var res []usecase.FindAllTagOutput
	for _, tag := range tags {
		res = append(res, usecase.FindAllTagOutput{
			Id:   tag.ID(),
			Name: tag.Name(),
		})
	}
	return res
}
