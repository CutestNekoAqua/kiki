package provider

import "gitea.code-infection.com/efertone/kiki/pkg/model"

// Entry is a proxy struct to pass entry data around.
type Entry struct {
	ID      string
	Title   string
	Link    string
	Content string
}

// ToModel returns with the model representation
// of an entry.
func (e Entry) ToModel() *model.Entry {
	return &model.Entry{
		EntryID: e.ID,
		Title:   e.Title,
		Link:    e.Link,
		Content: e.Content,
	}
}
