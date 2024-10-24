package dto

import "github.com/Reimei1213/lab/graphql-relay/domain/entity"

type ListUsersResponse struct {
	Users           entity.Users
	HasNextPage     bool
	HasPreviousPage bool
}
