package feed

import (
	"fmt"
)

// UserDoesNotExistError occues when user already does not exist
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
