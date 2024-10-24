//go:generate go run github.com/99designs/gqlgen generate
package controllers

import (
	"context"

	"github.com/Reimei1213/lab/graphql-relay/interface/inputport"
	"github.com/Reimei1213/lab/graphql-relay/pkg/db"
	"github.com/Reimei1213/lab/graphql-relay/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserInputport inputport.User
	CardInputport inputport.Card
}

func NewResolver(ctx context.Context) *Resolver {
	userUsecase := usecase.NewUser(db.NewUser(ctx))
	cardUsecase := usecase.NewCard(db.NewCard(ctx))
	userInputport := inputport.NewUser(userUsecase)
	cardInputport := inputport.NewCard(cardUsecase)
	return &Resolver{
		UserInputport: userInputport,
		CardInputport: cardInputport,
	}
}
