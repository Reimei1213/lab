package graph

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type NodeType string

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
