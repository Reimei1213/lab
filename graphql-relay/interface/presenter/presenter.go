package presenter

import "github.com/Reimei1213/lab/graphql-relay/pkg/graph/model"

type NodeType string

type node interface {
	GetID() string
	GetNodeType() NodeType
	ToGraphqlNode() model.Node
}

func ToGraphqlModels(nodes []node) []model.Node {
	res := make([]model.Node, 0, len(nodes))
	for _, n := range nodes {
		res = append(res, n.ToGraphqlNode())
	}
	return res
}
