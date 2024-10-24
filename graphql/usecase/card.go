package usecase

import (
	"context"

	"github.com/Reimei1213/lab/graphql/domain/entity"
	"github.com/Reimei1213/lab/graphql/domain/repository"
)

type Card interface {
	List(ctx context.Context) (entity.Cards, error)
	Get(ctx context.Context, id string) (*entity.Card, error)
	Create(ctx context.Context, title string, userID *string) (string, error)
}

type CardImpl struct {
	cardRepository repository.Card
	generateID     func() (string, error)
}

var _ Card = (*CardImpl)(nil)

func NewCard(cardRepository repository.Card) *CardImpl {
	return &CardImpl{
		cardRepository: cardRepository,
		generateID:     generateID,
	}
}

func (u *CardImpl) List(ctx context.Context) (entity.Cards, error) {
	return nil, nil
}

func (u *CardImpl) Get(ctx context.Context, id string) (*entity.Card, error) {
	return u.cardRepository.Get(ctx, id)
}

func (u *CardImpl) Create(ctx context.Context, title string, userID *string) (string, error) {
	id, err := u.generateID()
	if err != nil {
		return "", err
	}

	card := &entity.Card{
		ID:     id,
		Title:  title,
		UserID: userID,
		Status: entity.CardStatusTodo,
	}

	if err := u.cardRepository.Create(ctx, card); err != nil {
		return "", err
	}
	return id, nil
}
