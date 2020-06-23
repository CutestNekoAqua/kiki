package provider

import (
	"log"
)

// Interface is a common interface for providers like RSS.
type Interface interface {
	Name() string
	Parse(body []byte) (entries []*Entry)
}

// NewProviderByName creates a new Provider based on its name.
func NewProviderByName(name string) Interface {
	switch name {
	case "atom":
		return &Atom{}
	case "rss":
		return &RSS{}
	case "rdf":
		return &RDF{}
	default:
		log.Fatalf("Unknown Provider: %s\n", name)
	}

	return nil
}
