package repository

import (
	"context"

	"github.com/Reimei1213/lab/graphql/domain/entity"
)

type User interface {
	List(ctx context.Context, first *int, after *string, last *int, before *string) (entity.Users, error)
	Get(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
}
