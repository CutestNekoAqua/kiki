package provider

import "log"

// Interface is a common interface for providers like Misskey.
type Interface interface {
	Name() string
	Publish(content string) error
}

// NewTokenProviderByName creates a new Provider with token auth
// based on its name.
func NewTokenProviderByName(name, baseURL, token string) Interface {
	switch name {
	case "misskey":
		return NewMisskey(baseURL, token)
	default:
		log.Fatalf("Unknown provider: %s\n", name)
	}
}
