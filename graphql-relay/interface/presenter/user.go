package presenter

import (
	"github.com/Reimei1213/lab/graphql-relay/domain/entity"
	"github.com/Reimei1213/lab/graphql-relay/pkg/graph"
	"github.com/Reimei1213/lab/graphql-relay/pkg/graph/model"
)

const NodeTypeUser NodeType = "users"

type User struct {
	*entity.User
}

var _ Node = (*User)(nil)

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetNodeType() NodeType {
	return NodeTypeUser
}

func (u *User) ToGraphqlModel() *model.User {
	if u == nil {
		return nil
	}
	return &model.User{
		ID:   graph.EncodeGraphqlID(u.GetNodeType(), u.ID),
		Name: u.Name,
	}
}

func (u *User) ToGraphqlNode() model.Node {
	return model.Node(u.ToGraphqlModel())
}

func ToUser(u *entity.User) *User {
	return &User{u}
}

func ToUsers(us entity.Users) []*User {
	res := make([]*User, 0, len(us))
	for _, u := range us {
		res = append(res, ToUser(u))
	}
	return res
}
