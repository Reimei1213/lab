package repository

import (
	"context"

	"github.com/Reimei1213/lab/graphql/domain/entity"
)

type Card interface {
	Get(ctx context.Context, id string) (*entity.Card, error)
	Create(ctx context.Context, card *entity.Card) error
}
