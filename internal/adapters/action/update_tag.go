package action

import (
	"context"
	tagsv1 "neuro-most/tags-service/gen/go/tags"
	"neuro-most/tags-service/internal/usecase"
)

type UpdateTagAction struct {
	uc usecase.UpdateTagUseCase
}

func NewUpdateTagAction(uc usecase.UpdateTagUseCase) UpdateTagAction {
	return UpdateTagAction{uc: uc}
}

func (a UpdateTagAction) Execute(ctx context.Context, input *tagsv1.UpdateTagsRequest) error {
	return a.uc.Execute(ctx, usecase.UpdateTagInput{Id: input.Id, Name: input.Name})
}
