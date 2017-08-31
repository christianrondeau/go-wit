// Copyright (c) 2014 Jason Goecke
// messages_test.go

package wit

import (
	"os"
	"testing"
)

var msgID string

func TestWitPostAudioMessage(t *testing.T) {
	client := NewClient(os.Getenv("WIT_ACCESS_TOKEN"))
	request := &MessageRequest{}
	request.File = "./_audio_sample/helloWorld.wav"
	request.ContentType = "audio/wav"
	message, err := client.AudioMessage(request)
	if err != nil || message == nil {
		t.Error(err)
	} else {
		if message.Text != "hello world" {
			t.Error("Audio POST did not work properly")
		}
	}
}

func TestWitPostAudioContentsMessage(t *testing.T) {
	client := NewClient(os.Getenv("WIT_ACCESS_TOKEN"))
	file, err := os.Open("./_audio_sample/helloWorld.wav")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()
	stats, statsErr := file.Stat()
	if statsErr != nil {
		t.Error(err)
	}
	size := stats.Size()
	data := make([]byte, size)
	file.Read(data)
	request := &MessageRequest{}
	request.FileContents = data
	request.ContentType = "audio/wav"
	message, err := client.AudioMessage(request)
	if err != nil || message == nil {
		t.Error(err)
	} else {
		if message.Text != "hello world" {
			t.Error("Audio POST did not work properly")
		}
	}
}
