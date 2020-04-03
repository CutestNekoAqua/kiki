package misskey

import (
	"encoding/json"
)

// BaseRequest is the base request
type BaseRequest struct {
	APIToken string
	Path     string
	Request  interface{}
}

// ToJSON returns with the JSON []byte representation of the request
func (r BaseRequest) ToJSON() []byte {
	requestBody, _ := json.Marshal(r.Request)

	var repack map[string]interface{}
	json.Unmarshal(requestBody, &repack)
	repack["i"] = r.APIToken

	requestBody, _ = json.Marshal(repack)

	return requestBody
}

// SetAPIToken stores the API key for a Misskey User
func (r *BaseRequest) SetAPIToken(token string) {
	r.APIToken = token
}
