package feed

import (
	"fmt"
)

// UserDoesNotExistError occues when user does not exist
// but we want to assign a feed to it.
type UserDoesNotExistError struct {
	Name string
}

func (e UserDoesNotExistError) Error() string {
	return fmt.Sprintf(
		"User does not exist: %s",
		e.Name,
	)
}

// DoesNotExistError occues when feed does not exist.
type DoesNotExistError struct {
	ID uint
}

func (e DoesNotExistError) Error() string {
	return fmt.Sprintf(
		"Feed does not exist: %d",
		e.ID,
	)
}
