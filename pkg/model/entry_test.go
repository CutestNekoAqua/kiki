package model_test

import (
	"fmt"
	"testing"

	"gitea.code-infection.com/efertone/kiki/test"

	"gitea.code-infection.com/efertone/kiki/pkg/model"
)

type testCase struct {
	name        string
	fixtureFile string
}

func TestEntry_Excerpt(t *testing.T) {
	tests := []testCase{
		{
			name:        "Short content",
			fixtureFile: "entry/short_content",
		},
		{
			name:        "Long content",
			fixtureFile: "entry/long_content",
		},
		{
			name:        "Long content with divider",
			fixtureFile: "entry/long_content_divider",
		},
		{
			name:        "Long content without punctuations",
			fixtureFile: "entry/long_content_no_punctuations",
		},
		{
			name:        "Long content without punctuations and newline",
			fixtureFile: "entry/long_content_no_punctuations_no_newline",
		},
		{
			name:        "Long content big block",
			fixtureFile: "entry/long_content_big_block",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			e := model.Entry{
				Content: test.Must(test.LoadFixture(tt.fixtureFile)),
			}
			want := test.Must(test.LoadFixture(fmt.Sprintf("%s.expected", tt.fixtureFile)))
			if got := e.Excerpt(); got != want {
				t.Errorf("Excerpt() = %v, want %v", got, want)
			}
		})
	}
}
