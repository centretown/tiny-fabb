package settings

import (
	"testing"
)

func TestLoadLocalSettings(t *testing.T) {
	err := LocalSettings.Load()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(LocalSettings)
}
