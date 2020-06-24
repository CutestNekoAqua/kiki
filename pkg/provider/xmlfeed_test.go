package provider_test

import (
	"testing"

	"gitea.code-infection.com/efertone/kiki/pkg/provider"
	"gitea.code-infection.com/efertone/kiki/test"
)

func TestXMLFeedTypeOf_Invalid(t *testing.T) {
	input := []byte(`this is not an xml file`)

	_, err := provider.XMLFeedTypeOf(input)
	if err == nil {
		t.Error("Expected error, nothing happened")
		return
	}

	expectedErrorMessage := "EOF"
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error = %s; got = %s", expectedErrorMessage, err.Error())
	}
}

func TestXMLFeedTypeOf_Unknown(t *testing.T) {
	input := []byte(`<?xml version="1.0" encoding="utf-8" standalone="yes"?>
	<something version="2.0" xmlns:something="http://www.w3.org/2005/Something"></something>`)

	_, err := provider.XMLFeedTypeOf(input)
	if err == nil {
		t.Error("Expected error, nothing happened")
		return
	}

	expectedError := provider.UnknownXMLFeedType{Type: "something"}
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error = %s; got = %s", expectedError.Error(), err.Error())
	}
}

func TestXMLFeedTypeOf_Atom(t *testing.T) {
	expected := provider.AtomName
	input := []byte(test.Must(test.LoadFixture("atom.xml")))

	name, err := provider.XMLFeedTypeOf(input)
	if err != nil {
		t.Errorf("Unexpected error = %s", err)
		return
	}

	if name != expected {
		t.Errorf("Expected error = %s; got = %s", expected, name)
	}
}

func TestXMLFeedTypeOf_RSS(t *testing.T) {
	expected := provider.RSSName
	input := []byte(test.Must(test.LoadFixture("rss.xml")))

	name, err := provider.XMLFeedTypeOf(input)
	if err != nil {
		t.Errorf("Unexpected error = %s", err)
		return
	}

	if name != expected {
		t.Errorf("Expected error = %s; got = %s", expected, name)
	}
}

func TestXMLFeedTypeOf_RDF(t *testing.T) {
	expected := provider.RDFName
	input := []byte(test.Must(test.LoadFixture("rdf.xml")))

	name, err := provider.XMLFeedTypeOf(input)
	if err != nil {
		t.Errorf("Unexpected error = %s", err)
		return
	}

	if name != expected {
		t.Errorf("Expected error = %s; got = %s", expected, name)
	}
}
