// Copyright (c) 2014 Jason Goecke
// messages_test.go

package wit

import (
	"os"
	"testing"
)

func TestWitMessage(t *testing.T) {
	client := NewClient(os.Getenv("WIT_ACCESS_TOKEN"))
	message, err := client.Message("Hello World")
	if err != nil || message == nil {
		t.Error(err)
		return
	}

	if message.Text != "Hello World" {
		t.Errorf("Received '%v' instead of 'Hello World'", message.Text)
	}
}
