package presenter

import (
	"github.com/Reimei1213/lab/graphql-relay/domain/entity"
	"github.com/Reimei1213/lab/graphql-relay/pkg/graph/model"
)

const NodeTypeCard NodeType = "cards"

type Card struct {
	*entity.Card
}

var _ node = (*Card)(nil)

func (c *Card) GetID() string {
	return c.ID
}

func (c *Card) GetNodeType() NodeType {
	return NodeTypeCard
}

func (c *Card) ToGraphqlModel() *model.Card {
	var userID string
	if c == nil {
		return nil
	}
	if c.UserID != nil {
		userID = EncodeGraphqlID(NodeTypeUser, *c.UserID)
	}
	return &model.Card{
		ID:           EncodeGraphqlID(c.GetNodeType(), c.ID),
		Title:        c.Title,
		Status:       NewCardStatus(c.Status),
		AssignedUser: &model.User{ID: userID},
	}
}

func (c *Card) ToGraphqlNode() model.Node {
	return model.Node(c.ToGraphqlModel())
}

func NewCardStatus(in entity.CardStatus) model.CardStatus {
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

func NewCard(c *entity.Card) *Card {
	return &Card{c}
}

func NewCardNodes(cs entity.Cards) []node {
	res := make([]node, 0, len(cs))
	for _, c := range cs {
		res = append(res, NewCard(c))
	}
	return res
}
