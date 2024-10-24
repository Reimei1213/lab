package db

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/Reimei1213/lab/graphql-relay/domain/entity"
	"github.com/Reimei1213/lab/graphql-relay/domain/repository"
)

type Card struct {
	client *gorm.DB
	now    time.Time
}

var _ repository.Card = (*Card)(nil)

func NewCard(ctx context.Context) *Card {
	db := NewDB(ctx)
	return &Card{
		client: db.client,
		now:    time.Now(),
	}
}

func (db *Card) Get(_ context.Context, id string) (*entity.Card, error) {
	var card *entity.Card
	if err := db.client.First(&card, "id = ? AND is_active = 1", id).Error; err != nil {
		return nil, err
	}
	return card, nil
}

func (db *Card) Create(_ context.Context, card *entity.Card) error {
	card.CreatedAt = db.now.Unix()
	card.UpdatedAt = db.now.Unix()

	if err := db.client.Create(card).Error; err != nil {
		return err
	}
	return nil
}
