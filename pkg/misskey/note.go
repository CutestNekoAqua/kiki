package misskey

// NoteCreateRequest represents a CreateNote request
type NoteCreateRequest struct {
	*BaseRequest
	Visibility        string   `json:"visibility"`
	VisibleUserIDs    []string `json:"visibleUserIds,omitempty"`
	Text              string   `json:"text,omitempty"`
	CW                string   `json:"cw,omitempty"`
	ViaMobile         bool     `json:"viaMobile"`
	LocalOnly         bool     `json:"localOnly"`
	NoExtractMentions bool     `json:"noExtractMentions"`
	NoExtractHashtags bool     `json:"noExtractHashtags"`
	NoExtractEmojis   bool     `json:"noExtractEmojis"`
	FileIDs           []string `json:"fileIds,omitempty"`
	ReplyID           string   `json:"replyId,omitempty"`
	RenoteID          string   `json:"renoteId,omitempty"`
	Poll              *Poll    `json:"poll,omitempty"`
}
