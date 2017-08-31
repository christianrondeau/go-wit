package _examples

import (
	"github.com/christianrondeau/go-wit"
	"log"
	"os"
)

func main() {
	client := wit.NewClient(os.Getenv("WIT_ACCESS_TOKEN"))

	// Show a Wit builtin entity
	entity, _ := client.Entity("wit$age_of_person")
	log.Println(entity.Name)

	// Show a custom entity
	entity, _ = client.Entity("68e5fbfb-1839-422a-8e52-28798344b2ad$favorite_city")
	log.Println(entity.Name)
}
