package action

import (
	"context"
	tagsv1 "neuro-most/tags-service/gen/go/tags"
	"neuro-most/tags-service/internal/usecase"
)

type FindAllTagAction struct {
	uc usecase.FindAllTag
}

func NewFindAllTagAction(uc usecase.FindAllTag) FindAllTagAction {
	return FindAllTagAction{uc: uc}
}

func (a FindAllTagAction) Execute(ctx context.Context, input *tagsv1.GetTagsFeedRequest) (*tagsv1.GetTagsFeedResponse, error) {
	var usecaseInput usecase.FindAllTagInput
	usecaseInput.Page = int64(input.Page)
	usecaseInput.PageSize = int64(input.PageSize)
	tags, total, err := a.uc.Execute(ctx, usecaseInput)
	if err != nil {
		return nil, err
	}
	var tagsResponse []*tagsv1.Tags
	for _, tag := range tags {
		tagsResponse = append(tagsResponse, &tagsv1.Tags{
			Id:   tag.Id,
			Name: tag.Name,
		})
	}
	return &tagsv1.GetTagsFeedResponse{
		Tags:  tagsResponse,
		Total: int32(total),
	}, nil
}
