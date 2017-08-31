package _examples

import (
	"github.com/christianrondeau/go-wit"
	"log"
	"os"
)

func PrintIntents() {
	client := wit.NewClient(os.Getenv("WIT_ACCESS_TOKEN"))

	result, _ := client.Intents()
	log.Println(result)
}
