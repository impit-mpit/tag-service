package action

import (
	"context"
	tagsv1 "neuro-most/tags-service/gen/go/tags"
	"neuro-most/tags-service/internal/usecase"
)

type GetByIDTagAction struct {
	uc usecase.GetByIdTagUseCase
}

func NewGetByIDTagAction(uc usecase.GetByIdTagUseCase) GetByIDTagAction {
	return GetByIDTagAction{uc: uc}
}

func (a GetByIDTagAction) Execute(ctx context.Context, input *tagsv1.GetTagsByIdRequest) (*tagsv1.Tags, error) {
	tag, err := a.uc.Execute(ctx, usecase.GetByIdTagInput{Id: input.Id})
	if err != nil {
		return nil, err
	}
	return &tagsv1.Tags{
		Id:   tag.Id,
		Name: tag.Name,
	}, nil
}
