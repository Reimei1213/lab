package presenter

import (
	"github.com/Reimei1213/lab/graphql-relay/pkg/graph"
	"github.com/Reimei1213/lab/graphql-relay/pkg/graph/model"
)

type NodeType = graph.NodeType

type node interface {
	GetID() string
	GetNodeType() NodeType
	ToGraphqlNode() model.Node
}

func toGraphqlModels[T node](nodes []T) []model.Node {
	res := make([]model.Node, 0, len(nodes))
	for _, n := range nodes {
		res = append(res, n.ToGraphqlNode())
	}
	return res
}

func NewConnection[T node](nodes []T, hasNextPage, hasPreviousPage bool) *model.Connection {
	return &model.Connection{
		Edges:    newEdges(nodes),
		Nodes:    toGraphqlModels(nodes),
		PageInfo: newPageInfo(nodes, hasNextPage, hasPreviousPage),
	}
}

func newEdges[T node](nodes []T) []*model.Edge {
	res := make([]*model.Edge, 0, len(nodes))
	for _, n := range nodes {
		res = append(res, &model.Edge{
			Cursor: graph.EncodeGraphqlID(n.GetNodeType(), n.GetID()),
			Node:   n.ToGraphqlNode(),
		})
	}
	return res
}

func newPageInfo[T node](nodes []T, hasNextPage, hasPreviousPage bool) *model.PageInfo {
	if len(nodes) == 0 {
		return &model.PageInfo{
			HasPreviousPage: hasPreviousPage,
			HasNextPage:     hasNextPage,
		}
	}

	startNode := nodes[0]
	endNode := nodes[len(nodes)-1]
	startCursor := graph.EncodeGraphqlID(startNode.GetNodeType(), startNode.GetID())
	endCursor := graph.EncodeGraphqlID(endNode.GetNodeType(), endNode.GetID())
	return &model.PageInfo{
		HasPreviousPage: hasPreviousPage,
		StartCursor:     &startCursor,
		HasNextPage:     hasNextPage,
		EndCursor:       &endCursor,
	}
}
