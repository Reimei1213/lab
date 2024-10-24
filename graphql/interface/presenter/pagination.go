package presenter

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/Reimei1213/lab/graphql/pkg/graph/model"
)

const graphqlIDFormat = "%s:%s" // tableID:id

func EncodeGraphqlID(nodeType NodeType, id string) string {
	graphqlID := fmt.Sprintf(graphqlIDFormat, nodeType, id)
	return base64.StdEncoding.EncodeToString([]byte(graphqlID))
}

func DecodeGraphqlID(encodedGraphqlID string) (NodeType, string, error) {
	graphqlID, err := base64.StdEncoding.DecodeString(encodedGraphqlID)
	if err != nil {
		return "", "", err
	}

	result := strings.Split(string(graphqlID), ":")

	return NodeType(result[0]), result[1], nil
}

func NewConnection(nodes []node, hasNextPage, hasPreviousPage bool) *model.Connection {
	return &model.Connection{
		Edges:    NewEdges(nodes),
		Nodes:    ToGraphqlModels(nodes),
		PageInfo: NewPageInfo(nodes, hasNextPage, hasPreviousPage),
	}
}

func NewEdges(nodes []node) []*model.Edge {
	res := make([]*model.Edge, 0, len(nodes))
	for _, n := range nodes {
		res = append(res, &model.Edge{
			Cursor: EncodeGraphqlID(n.GetNodeType(), n.GetID()),
			Node:   n.ToGraphqlNode(),
		})
	}
	return res
}

func NewPageInfo(nodes []node, hasNextPage, hasPreviousPage bool) *model.PageInfo {
	if len(nodes) == 0 {
		return &model.PageInfo{
			HasPreviousPage: hasPreviousPage,
			HasNextPage:     hasNextPage,
		}
	}

	startNode := nodes[0]
	endNode := nodes[len(nodes)-1]
	startCursor := EncodeGraphqlID(startNode.GetNodeType(), startNode.GetID())
	endCursor := EncodeGraphqlID(endNode.GetNodeType(), endNode.GetID())
	return &model.PageInfo{
		HasPreviousPage: hasPreviousPage,
		StartCursor:     &startCursor,
		HasNextPage:     hasNextPage,
		EndCursor:       &endCursor,
	}
}
