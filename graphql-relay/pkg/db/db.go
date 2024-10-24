package db

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/Reimei1213/lab/graphql-relay/domain/entity"
)

type DB struct {
	client *gorm.DB
}

func NewDB(_ context.Context) *DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(
		&entity.User{},
		&entity.Card{},
	); err != nil {
		panic(err)
	}

	return &DB{client: db}
}

type listBaseQueryParams struct {
	first  *int
	after  *string
	last   *int
	before *string
	db     *gorm.DB
}

func listBaseQuery(_ context.Context, params *listBaseQueryParams) *gorm.DB {
	query := params.db.Where("is_active = ?", true).Order("id ASC")

	if params.after != nil {
		query = query.Where("id > ?", *params.after)
	}

	if params.before != nil {
		query = query.Where("id < ?", *params.before)
	}

	limit := 0
	if params.first != nil {
		limit = *params.first + 1
		query = query.Order("id ASC").Limit(limit)
	} else if params.last != nil {
		limit = *params.last + 1
		query = query.Order("id DESC").Limit(limit)
	}

	return query
}
