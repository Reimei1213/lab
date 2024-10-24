package db

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/Reimei1213/lab/graphql-relay/domain/entity"
	"github.com/Reimei1213/lab/graphql-relay/domain/repository"
)

type User struct {
	client *gorm.DB
	now    time.Time
}

var _ repository.User = (*User)(nil)

func NewUser(ctx context.Context) *User {
	db := NewDB(ctx)
	return &User{
		client: db.client,
		now:    time.Now(),
	}
}

func (db *User) List(ctx context.Context, first *int, after *string, last *int, before *string) (entity.Users, error) {
	var users entity.Users
	query := listBaseQuery(ctx, &listBaseQueryParams{
		first:  first,
		after:  after,
		last:   last,
		before: before,
		db:     db.client,
	})

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (db *User) Get(_ context.Context, id string) (*entity.User, error) {
	var user *entity.User
	if err := db.client.First(&user, "id = ? AND is_active = 1", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (db *User) Create(_ context.Context, user *entity.User) error {
	user.CreatedAt = db.now.Unix()
	user.UpdatedAt = db.now.Unix()

	if err := db.client.Create(user).Error; err != nil {
		return err
	}
	return nil
}
