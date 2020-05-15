package misskey

// Poll represents a Poll data structure for Misskey.
type Poll struct {
	Choices      []string `json:"choices,omitempty"`
	Multiple     bool     `json:"multiple,omitempty"`
	ExpiresAt    uint     `json:"expiresAt,omitempty"`
	ExpiredAfter uint     `json:"expiredAfter,omitempty"`
}
