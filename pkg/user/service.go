package user

import (
	"context"
	"fmt"
	"myproject/pkg/model"
)

type Service interface {
	Register(ctx context.Context, request model.UserRegisterRequest) error
	Listing(ctx context.Context) ([]model.UserRegisterRequest, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Register(ctx context.Context, request model.UserRegisterRequest) error {
	fmt.Println("this is in the service Register")

	return s.repo.Register(ctx, request)
}

// func (s *service) Listing(ctx context.Context) ([]model.Product, error) {

//		return s.repo.Listing(ctx)
//	}
func (s *service) Listing(ctx context.Context) ([]model.UserRegisterRequest, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return s.repo.Listing(ctx)
	}
}
