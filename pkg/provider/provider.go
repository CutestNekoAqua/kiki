package provider

// Interface is a common interface for providers like RSS.
type Interface interface {
	Name() string
	Parse(body []byte) (entries []*Entry)
}

// NewProviderByName creates a new Provider based on its name.
func NewProviderByName(name string) Interface {
	switch name {
	case AtomName:
		return &Atom{}
	case RSSName:
		return &RSS{}
	case RDFName:
		return &RDF{}
	}

	return nil
}
