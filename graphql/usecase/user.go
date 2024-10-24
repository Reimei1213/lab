package usecase

import (
	"context"

	"github.com/Reimei1213/lab/graphql/domain/entity"
	"github.com/Reimei1213/lab/graphql/domain/repository"
	"github.com/Reimei1213/lab/graphql/usecase/dto"
)

type User interface {
	List(ctx context.Context, first *int, after *string, last *int, before *string) (*dto.ListUsersResponse, error)
	Get(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, name string) (string, error)
}

type UserImpl struct {
	userRepository repository.User
	generateID     func() (string, error)
}

var _ User = (*UserImpl)(nil)

func NewUser(userRepository repository.User) *UserImpl {
	return &UserImpl{
		userRepository: userRepository,
		generateID:     generateID,
	}
}

func (u *UserImpl) List(ctx context.Context, first *int, after *string, last *int, before *string) (*dto.ListUsersResponse, error) {
	users, err := u.userRepository.List(ctx, first, after, last, before)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return &dto.ListUsersResponse{}, nil
	}

	if last != nil {
		// 逆順で取得しているので反転
		for i, j := 0, len(users)-1; i < j; i, j = i+1, j-1 {
			users[i], users[j] = users[j], users[i]
		}
	}

	users, hasNextPage, hasPreviousPage := evaluatePageNavigation(users, first, last)
	return &dto.ListUsersResponse{
		Users:           users,
		HasNextPage:     hasNextPage,
		HasPreviousPage: hasPreviousPage,
	}, nil
}

func (u *UserImpl) Get(ctx context.Context, id string) (*entity.User, error) {
	return u.userRepository.Get(ctx, id)
}

func (u *UserImpl) Create(ctx context.Context, name string) (string, error) {
	id, err := u.generateID()
	if err != nil {
		return "", err
	}

	user := &entity.User{
		ID:   id,
		Name: name,
	}

	if err := u.userRepository.Create(ctx, user); err != nil {
		return "", err
	}

	return id, nil
}
