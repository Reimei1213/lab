package inputport

import (
	"context"

	"github.com/Reimei1213/lab/graphql/domain/entity"
	"github.com/Reimei1213/lab/graphql/pkg/graph/model"
	"github.com/Reimei1213/lab/graphql/usecase"
	"github.com/Reimei1213/lab/graphql/usecase/dto"
)

type User interface {
	List(ctx context.Context, first *int, after *string, last *int, before *string) (*dto.ListUsersResponse, error)
	Get(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, input *model.AddUserInput) (string, error)
}

type UserImpl struct {
	usecase usecase.User
}

var _ User = (*UserImpl)(nil)

func NewUser(usecase usecase.User) *UserImpl {
	return &UserImpl{
		usecase: usecase,
	}
}

func (i *UserImpl) List(ctx context.Context, first *int, after *string, last *int, before *string) (*dto.ListUsersResponse, error) {
	return i.usecase.List(ctx, first, after, last, before)
}

func (i *UserImpl) Get(ctx context.Context, id string) (*entity.User, error) {
	return i.usecase.Get(ctx, id)
}

func (i *UserImpl) Create(ctx context.Context, input *model.AddUserInput) (string, error) {
	return i.usecase.Create(ctx, input.Name)
}
