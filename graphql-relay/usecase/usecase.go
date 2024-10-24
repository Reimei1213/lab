package usecase

import "github.com/google/uuid"

func generateID() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func evaluatePageNavigation[T any, Slice ~[]T](s Slice, first, last *int) (Slice, bool, bool) {
	hasNextPage := false
	hasPreviousPage := false
	if first != nil && len(s) > *first {
		hasNextPage = true
		s = s[:*first]
	} else if last != nil && len(s) > *last {
		hasPreviousPage = true
		s = s[:*last]
	}
	return s, hasNextPage, hasPreviousPage
}
