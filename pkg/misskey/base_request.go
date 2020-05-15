package misskey

import (
	"encoding/json"
	"log"
)

// BaseRequest is the base request.
type BaseRequest struct {
	APIToken string
	Path     string
	Request  interface{}
}

// ToJSON returns with the JSON []byte representation of the request.
func (r BaseRequest) ToJSON() []byte {
	requestBody, _ := json.Marshal(r.Request)

	var repack map[string]interface{}

	if err := json.Unmarshal(requestBody, &repack); err != nil {
		log.Fatalf("Lethal damage: %s\n", err)
	}

	repack["i"] = r.APIToken

	requestBody, _ = json.Marshal(repack)

	return requestBody
}

// SetAPIToken stores the API key for a Misskey User.
func (r *BaseRequest) SetAPIToken(token string) {
	r.APIToken = token
}
