package misskey

import (
	"encoding/json"
)

type BaseRequest struct {
	APIToken string
	Path     string
	Request  interface{}
}

func (r BaseRequest) ToJSON() []byte {
	requestBody, _ := json.Marshal(r.Request)

	var repack map[string]interface{}
	json.Unmarshal(requestBody, &repack)
	repack["i"] = r.APIToken

	requestBody, _ = json.Marshal(repack)

	return requestBody
}

func (r *BaseRequest) SetAPIToken(token string) {
	r.APIToken = token
}
