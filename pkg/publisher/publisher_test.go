package publisher_test

import (
	"testing"

	"gitea.code-infection.com/efertone/kiki/pkg/publisher"
)

func TestNewTokenPublisherByName_Misskey(t *testing.T) {
	pub := publisher.NewTokenPublisherByName("misskey", "url", "token")
	if pub == nil {
		t.Error("Expected Misskey Publisher, got nothing")
		return
	}

	if pub.Name() != "misskey" {
		t.Errorf("Expected Name() = misskey, got %s", pub.Name())
	}
}

func TestNewTokenPublisherByName_Unknown(t *testing.T) {
	pub := publisher.NewTokenPublisherByName("unknown", "url", "token")
	if pub != nil {
		t.Errorf("Unexpected Publisher: %s", pub.Name())
	}
}
