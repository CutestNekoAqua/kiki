package misskey

type Poll struct {
	Choices      []string `json:"choices,omitempty"`
	Multiple     bool     `json:"multiple,omitempty"`
	ExpiresAt    uint     `json:"expiresAt,omitempty"`
	ExpiredAfter uint     `json:"expiredAfter,omitempty"`
}
