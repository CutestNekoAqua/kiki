package hashtag

import (
	"fmt"
)

// FeedDoesNotExistError occues when user already does not exist
// but we want to assign a feed to it.
type FeedDoesNotExistError struct {
	ID uint
}

func (e FeedDoesNotExistError) Error() string {
	return fmt.Sprintf(
		"Feed does not exist: %d",
		e.ID,
	)
}
