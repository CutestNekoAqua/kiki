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
	// Request timeout in seconds.
	RequestTimout = 10
)

// NewClient creates a new Misskey Client.
func NewClient(baseURL, token string) *Client {
	return &Client{Token: token, BaseURL: baseURL}
}

func (c Client) url(path string) string {
	return fmt.Sprintf("%s/api%s", c.BaseURL, path)
}

func (c Client) sendRequest(request *BaseRequest) bool {
	request.SetAPIToken(c.Token)
	requestBody := request.ToJSON()

	req, err := http.NewRequest("POST", c.url(request.Path), bytes.NewBuffer(requestBody))
	if err != nil {
		log.Printf("[Misskey] Error reading request: %s\n", err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Kiki, News Delivery Service")

	client := &http.Client{Timeout: time.Second * RequestTimout}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("[Misskey] Error reading response: %s\n", err)
		return false
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("[Misskey] Error reading body: %s\n", err)
		return false
	}

	if string(body) == "Not Found" {
		return false
	}

	if resp.StatusCode == http.StatusOK {
		return true
	}

	var errorWrapper struct {
		Error json.RawMessage `json:"error"`
	}

	err = json.Unmarshal(body, &errorWrapper)
	if err != nil {
		log.Println(err)
		return false
	}

	var requestError ErrorResponse
	if err := json.Unmarshal(errorWrapper.Error, &requestError); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}

	log.Printf("[Misskey] <%s> %s -> %s", requestError.Code, requestError.Info.Param, requestError.Info.Reason)

	return false
}

// CreateNote sends a request to the Misskey server to create a note.
func (c *Client) CreateNote(content string) bool {
	request := &NoteCreateRequest{
		Visibility: "public",
		Text:       content,
	}

	return c.sendRequest(&BaseRequest{Request: request, Path: "/notes/create"})
}
