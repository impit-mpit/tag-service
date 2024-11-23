package entities

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrTagsNotFound = status.New(codes.NotFound, "tags not found").Err()
	ErrorTagsCreate = status.New(codes.Internal, "error create tags").Err()
	ErrorTagsUpdate = status.New(codes.Internal, "error update tags").Err()
	ErrorTagsDelete = status.New(codes.Internal, "error delete tags").Err()
	ErrorTagsFetch  = status.New(codes.Internal, "error fetch news").Err()
)

type (
	TagsRepo interface {
		Create(ctx context.Context, tag Tags) error
		Update(ctx context.Context, tag Tags) error
		Delete(ctx context.Context, tag Tags) error
		GetByID(ctx context.Context, id int64) (Tags, error)
		Fetch(ctx context.Context, page, pageSize int64) ([]Tags, int64, error)
	}

	Tags struct {
		id   int64
		name string
	}
)

func NewTag(
	id int64,
	name string,
) Tags {
	return Tags{
		id:   id,
		name: name,
	}
}

func NewTagCreate(
	name string,
) Tags {
	return Tags{
		name: name,
	}
}

func (t Tags) ID() int64 {
	return t.id
}

func (t Tags) Name() string {
	return t.name
}

func (t *Tags) SetID(id int64) {
	t.id = id
}

func (t *Tags) SetName(name string) {
	t.name = name
}
