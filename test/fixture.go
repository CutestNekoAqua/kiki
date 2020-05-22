package test

import (
	"fmt"
	"io/ioutil"
	"os"
)

func LoadFixture(name string) (string, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s", name))
	if err != nil {
		path, _ := os.Getwd()
		return "", FixtureFileNotFound{path, name}
	}

	return string(content), nil
}

func Must(content string, err error) string {
	if err != nil {
		panic(err)
	}

	return content
}
