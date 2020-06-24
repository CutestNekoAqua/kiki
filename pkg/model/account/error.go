package account

import (
	"fmt"

	"gitea.code-infection.com/efertone/kiki/pkg/model"
)

// UserAlreadyExistsError occues when user already exists
// but we want to add it as a new one.
type UserAlreadyExistsError struct {
	Account model.Account
}

func (e UserAlreadyExistsError) Error() string {
	return fmt.Sprintf(
		"User already exists: %s -> %s -> %s",
		e.Account.Name,
		e.Account.Publisher,
		e.Account.BaseURL,
	)
}
