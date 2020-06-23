package provider

import (
	"io/ioutil"
	"net/http"
)

// Download fetches the content of an URI and returns with a parsed []mode.Entry.
func Download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	return body, nil
}
