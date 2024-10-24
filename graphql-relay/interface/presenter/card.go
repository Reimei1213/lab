package presenter

import (
	"github.com/Reimei1213/lab/graphql-relay/domain/entity"
	"github.com/Reimei1213/lab/graphql-relay/pkg/graph/model"
	"github.com/Reimei1213/lab/graphql-relay/pkg/graph/pagination"
)

const NodeTypeCard pagination.NodeType = "cards"

type Card struct {
	*entity.Card
}

var _ pagination.Node = (*Card)(nil)

func (c *Card) GetID() string {
	return c.ID
}

func (c *Card) GetNodeType() pagination.NodeType {
	return NodeTypeCard
}

func (c *Card) ToGraphqlModel() *model.Card {
	var userID string
	if c == nil {
		return nil
	}
	if c.UserID != nil {
		userID = pagination.EncodeGraphqlID(NodeTypeUser, *c.UserID)
	}
	return &model.Card{
		ID:           pagination.EncodeGraphqlID(c.GetNodeType(), c.ID),
		Title:        c.Title,
		Status:       ToCardStatus(c.Status),
		AssignedUser: &model.User{ID: userID},
	}
}

func (c *Card) ToGraphqlNode() model.Node {
	return model.Node(c.ToGraphqlModel())
}

func ToCardStatus(in entity.CardStatus) model.CardStatus {
	switch in {
	case entity.CardStatusTodo:
		return model.CardStatusTodo
	case entity.CardStatusDoing:
		return model.CardStatusDoing
	case entity.CardStatusDone:
		return model.CardStatusDone
	default:
		return model.CardStatusTodo
	}
}

func ToCard(c *entity.Card) *Card {
	return &Card{c}
}

func ToCards(cs []*entity.Card) []*Card {
	res := make([]*Card, 0, len(cs))
	for _, c := range cs {
		res = append(res, ToCard(c))
	}
	return res
}
