package publisher_test

import (
	"testing"

	"gitea.code-infection.com/efertone/kiki/pkg/publisher"
)

type MockMisskeyClient struct {
}

func (m MockMisskeyClient) CreateNote(content string) error {
	return nil
}

func TestMisskeyPublisher(t *testing.T) {
	pub := publisher.NewMisskey("https://localhost", "token")

	pub.Client = &MockMisskeyClient{}

	if pub.Name() != "misskey" {
		t.Errorf("publisher.Name(): expected = misskey; got = %s", pub.Name())
	}

	err := pub.Publish("test")
	if err != nil {
		t.Errorf("Unexpedted error: %s", err)
	}
}
