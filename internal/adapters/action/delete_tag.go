package action

import (
	"context"
	tagsv1 "neuro-most/tags-service/gen/go/tags"
	"neuro-most/tags-service/internal/usecase"
)

type DeleteTagAction struct {
	uc usecase.DeleteTagUseCase
}

func NewDeleteTagAction(uc usecase.DeleteTagUseCase) DeleteTagAction {
	return DeleteTagAction{uc: uc}
}

func (a DeleteTagAction) Execute(ctx context.Context, input *tagsv1.DeleteTagsRequest) error {
	return a.uc.Execute(ctx, usecase.DeleteTagInput{Id: input.Id})
}
