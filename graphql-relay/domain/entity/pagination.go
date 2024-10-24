package entity

type PageInfo struct {
	HasPreviousPage bool
	StartCursor     *string
	HasNextPage     bool
	EndCursor       *string
}
