package dto

import "github.com/Reimei1213/lab/graphql/domain/entity"

type ListUsersResponse struct {
	Users           entity.Users
	HasNextPage     bool
	HasPreviousPage bool
}
