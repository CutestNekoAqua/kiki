package provider

import "gitea.code-infection.com/efertone/kiki/pkg/misskey"

// Misskey provider.
type Misskey struct {
	Client *misskey.Client
}

const (
	misskeyName = "misskey"
)

// NewMisskey creates a new Misskey provider.
func NewMisskey(baseURL, token string) *Misskey {
	return &Misskey{
		Client: misskey.NewClient(baseURL, token),
	}
}

// Name returns with the name of the Provider.
func (m *Misskey) Name() string {
	return misskeyName
}

// Publish simply publishes a Note.
func (m *Misskey) Publish(content string) error {
	return m.Client.CreateNote(content)
}
