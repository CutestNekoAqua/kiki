// nolint:dupl
package provider_test

import (
	"testing"

	"gitea.code-infection.com/efertone/kiki/pkg/provider"
	"gitea.code-infection.com/efertone/kiki/test"
)

func TestAtom_Parse_Invalid(t *testing.T) {
	prov := provider.NewProviderByName(provider.AtomName)

	input := []byte(`this is not an xml file`)
	entries := prov.Parse(input)

	if len(entries) > 0 {
		t.Errorf("Expected 0 enties, got %d", len(entries))
	}
}

func TestAtom_Parse_Valid(t *testing.T) {
	prov := provider.NewProviderByName(provider.AtomName)

	input := []byte(test.Must(test.LoadFixture("atom.xml")))
	entries := prov.Parse(input)

	if len(entries) == 0 {
		t.Errorf("Expected enties, got none of them")
		return
	}

	first := entries[0]

	var expected string

	expected = "tag:eu.finalfantasyxiv.com,2020:/pr/blog//7.2973"
	if first.ID != expected {
		t.Errorf("ID expected = %s; got = %s", expected, first.ID)
	}

	expected = "Out of the Shadows and Into Your Ears"
	if first.Title != expected {
		t.Errorf("Title expected = %s; got = %s", expected, first.Title)
	}

	expected = "https://eu.finalfantasyxiv.com/pr/blog/002973.html"
	if first.Link != expected {
		t.Errorf("Link expected = %s; got = %s", expected, first.Link)
	}

	expected = "Greetings everyone, Zhexos here."

	excerpt := first.ToModel().Excerpt()[:len(expected)]
	if excerpt != expected {
		t.Errorf("Excerpt expected = %s; got = %s", expected, excerpt)
	}
}
