package publisher

import "log"

// Interface is a common interface for publishers like Misskey.
type Interface interface {
	Name() string
	Publish(content string) error
}

// NewTokenPublisherByName creates a new Publisher with token auth
// based on its name.
func NewTokenPublisherByName(name, baseURL, token string) Interface {
	switch name {
	case "misskey":
		return NewMisskey(baseURL, token)
	default:
		log.Fatalf("Unknown Publisher: %s\n", name)
	}

	return nil
}
