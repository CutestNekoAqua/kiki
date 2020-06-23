package misskey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Client is the main Misskey client struct.
type Client struct {
	BaseURL string
	Token   string
}

const (
	// RequestTimout is the timeout of a request in seconds.
	RequestTimout = 10
)

// NewClient creates a new Misskey Client.
func NewClient(baseURL, token string) *Client {
	return &Client{Token: token, BaseURL: baseURL}
}

func (c Client) url(path string) string {
	return fmt.Sprintf("%s/api%s", c.BaseURL, path)
}

func (c Client) sendRequest(request *BaseRequest) error {
	request.SetAPIToken(c.Token)
	requestBody := request.ToJSON()

	req, err := http.NewRequest("POST", c.url(request.Path), bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Kiki, News Delivery Service")

	client := &http.Client{Timeout: time.Second * RequestTimout}
	resp, err := client.Do(req)

	if err != nil {
		return RequestError{Message: ResponseReadError, Origin: err}
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return RequestError{Message: ResponseReadBodyError, Origin: err}
	}

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	var errorWrapper errorResponse

	err = json.Unmarshal(body, &errorWrapper)
	if err != nil {
		return RequestError{Message: ErrorResponseParseError, Origin: err}
	}

	var errorResponse ErrorResponse
	if err := json.Unmarshal(errorWrapper.Error, &errorResponse); err != nil {
		return RequestError{Message: ErrorResponseParseError, Origin: err}
	}

	log.Printf("[Misskey] <%s> %s -> %s", errorResponse.Code, errorResponse.Info.Param, errorResponse.Info.Reason)

	return UnknownError{Response: errorResponse}
}

// CreateNote sends a request to the Misskey server to create a note.
func (c *Client) CreateNote(content string) error {
	request := &NoteCreateRequest{
		Visibility: "public",
		Text:       content,
	}

	return c.sendRequest(&BaseRequest{Request: request, Path: "/notes/create"})
}
