package provider_test

import (
	"testing"

	"gitea.code-infection.com/efertone/kiki/pkg/provider"
)

func TestNewProviderByName(t *testing.T) {
	checkProvider(
		t,
		provider.NewProviderByName(provider.AtomName),
		provider.AtomName,
	)

	checkProvider(
		t,
		provider.NewProviderByName(provider.RSSName),
		provider.RSSName,
	)

	checkProvider(
		t,
		provider.NewProviderByName(provider.RDFName),
		provider.RDFName,
	)

	prov := provider.NewProviderByName("Unknown")
	if prov != nil {
		t.Errorf("Unexpected provider: %s", prov.Name())
	}
}

func checkProvider(t *testing.T, prov provider.Interface, name string) {
	if prov == nil {
		t.Errorf("Expected %s Provider, got nothing", name)
		return
	}

	if prov.Name() != name {
		t.Errorf("Expected Provider Name = %s, got = %s", name, prov.Name())
		return
	}
}
