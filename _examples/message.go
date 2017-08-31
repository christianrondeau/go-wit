package main

import (
	"github.com/christianrondeau/go-wit"
	"log"
	"os"
)

func main() {
	// Initialize the client
	client := wit.NewClient(os.Getenv("WIT_ACCESS_TOKEN"))

	// Text processing
	result, _ := client.Message("Hello World")
	log.Println(result.Entities["intent"].value)

	// Speech processing
	request = &wit.MessageRequest{}
	request.File = "../_audio_sample/helloWorld.wav"
	request.ContentType = "audio/wav;rate=8000"
	result, _ = client.AudioMessage(request)
	log.Println(result.Entities["intent"].value)
}
