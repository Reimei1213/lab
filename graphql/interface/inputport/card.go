package inputport

import (
	"context"

	"github.com/Reimei1213/lab/graphql/domain/entity"
	"github.com/Reimei1213/lab/graphql/pkg/graph/model"
	"github.com/Reimei1213/lab/graphql/usecase"
)

type Card interface {
	Get(ctx context.Context, id string) (*entity.Card, error)
	Create(ctx context.Context, input *model.AddCardInput) (string, error)
}

type CardImpl struct {
	usecase usecase.Card
}

var _ Card = (*CardImpl)(nil)

func NewCard(usecaes usecase.Card) *CardImpl {
	return &CardImpl{
		usecase: usecaes,
	}
}

func (i *CardImpl) Get(ctx context.Context, id string) (*entity.Card, error) {
	return i.usecase.Get(ctx, id)
}

func (i *CardImpl) Create(ctx context.Context, input *model.AddCardInput) (string, error) {
	return i.usecase.Create(ctx, input.Title, input.UserID)
}
