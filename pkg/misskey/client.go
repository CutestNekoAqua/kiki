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

type Client struct {
	BaseURL string
	Token   string
}

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
		log.Printf("[Misskey] Error reading request. ", err)
		return false
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Kiki, News Delivery Service")
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("[Misskey] Error reading response. ", err)
		return false
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[Misskey] Error reading body. ", err)
		return false
	}

	if string(body) == "Not Found" {
		return false
	}

	if resp.StatusCode == 200 {
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
	json.Unmarshal(errorWrapper.Error, &requestError)
	log.Printf("[Misskey] <%s> %s -> %s", requestError.Code, requestError.Info.Param, requestError.Info.Reason)
	return false
}

func (c *Client) CreateNote(content string) bool {
	request := &NoteCreateRequest{
		Visibility: "public",
		Text:       content,
	}
	return c.sendRequest(&BaseRequest{Request: request, Path: "/notes/create"})
}