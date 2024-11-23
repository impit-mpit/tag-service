package action

import (
	"context"
	tagsv1 "neuro-most/tags-service/gen/go/tags"
	"neuro-most/tags-service/internal/usecase"
)

type CreateTagAction struct {
	uc usecase.CreateTagUseCase
}

func NewCreateTagAction(uc usecase.CreateTagUseCase) CreateTagAction {
	return CreateTagAction{uc: uc}
}

func (a CreateTagAction) Execute(ctx context.Context, input *tagsv1.CreateTagsRequest) error {
	return a.uc.Execute(ctx, usecase.CreateTagInput{Name: input.Name})
}
