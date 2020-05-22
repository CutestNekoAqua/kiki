package test_test

import (
	"testing"

	"gitea.code-infection.com/efertone/kiki/test"
)

func TestLoadFixture_found(t *testing.T) {
	data, err := test.LoadFixture("data")
	if err != nil {
		t.Errorf("TestLoadFixture_found should find the 'data' file")
	}

	expected := "yey\n"
	if data != expected {
		t.Errorf("Expected file content: %s, got: %s", expected, data)
	}
}

func TestLoadFixture_notFound(t *testing.T) {
	data, err := test.LoadFixture("nodata")
	if err == nil {
		t.Errorf("TestLoadFixture_notFound should not find the 'nodata' file")
	}

	expected := ""
	if data != expected {
		t.Errorf("Expected file content: %s, got: %s", expected, data)
	}
}

func TestMust_noError(t *testing.T) {
	input := "Yey"
	expected := "Yey"

	output := test.Must(input, nil)
	if output != expected {
		t.Errorf("Must() = %v, want %v", output, expected)
	}
}

func TestMust_hasError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestMust_hasError should have panicked!")
		}
	}()

	_ = test.Must("", test.FixtureFileNotFound{Path: "/path", Name: "file"})
}
