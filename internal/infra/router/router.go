package router

import (
	"context"
	"log"
	"net"
	tagsv1 "neuro-most/tags-service/gen/go/tags"
	"neuro-most/tags-service/internal/adapters/action"
	"neuro-most/tags-service/internal/adapters/presenter"
	"neuro-most/tags-service/internal/adapters/repo"
	"neuro-most/tags-service/internal/usecase"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Router struct {
	db repo.GSQL
	tagsv1.UnimplementedTagsServiceServer
}

func NewRouter(db repo.GSQL) Router {
	return Router{db: db}
}

func (r *Router) Listen() {
	port := ":3001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts = []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)
	tagsv1.RegisterTagsServiceServer(srv, r)

	log.Printf("Starting gRPC server on port %s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (r *Router) CreateTags(ctx context.Context, input *tagsv1.CreateTagsRequest) (*emptypb.Empty, error) {
	var (
		uc  = usecase.NewCreateTagInteractor(repo.NewTagsRepo(r.db))
		act = action.NewCreateTagAction(uc)
	)
	return nil, act.Execute(ctx, input)
}
func (r *Router) DeleteTags(ctx context.Context, input *tagsv1.DeleteTagsRequest) (*emptypb.Empty, error) {
	var (
		uc  = usecase.NewDeleteTagInteractor(repo.NewTagsRepo(r.db))
		act = action.NewDeleteTagAction(uc)
	)
	return nil, act.Execute(ctx, input)
}
func (r *Router) GetTagsById(ctx context.Context, input *tagsv1.GetTagsByIdRequest) (*tagsv1.Tags, error) {
	var (
		uc  = usecase.NewGetByIdTagInteractor(repo.NewTagsRepo(r.db), presenter.NewGetByIdPresenter())
		act = action.NewGetByIDTagAction(uc)
	)
	return act.Execute(ctx, input)
}
func (r *Router) GetTagsFeed(ctx context.Context, input *tagsv1.GetTagsFeedRequest) (*tagsv1.GetTagsFeedResponse, error) {
	var (
		uc  = usecase.NewFindAllTagInteractor(repo.NewTagsRepo(r.db), presenter.NewFindAllTagPresenter())
		act = action.NewFindAllTagAction(uc)
	)
	return act.Execute(ctx, input)
}
func (r *Router) UpdateTags(ctx context.Context, input *tagsv1.UpdateTagsRequest) (*emptypb.Empty, error) {
	var (
		uc  = usecase.NewUpdateTagInteractor(repo.NewTagsRepo(r.db))
		act = action.NewUpdateTagAction(uc)
	)
	return nil, act.Execute(ctx, input)
}
