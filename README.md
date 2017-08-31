#go-wit 

[![Build Status](https://api.travis-ci.org/christianrondeau/go-wit.svg)](http://travis-ci.org/christianrondeau/go-wit) [![GoDoc](https://godoc.org/github.com/christianrondeau/go-wit?status.svg)](http://godoc.org/github.com/christianrondeau/go-wit)

A Go library for the [Wit.ai](http://wit.ai) API for Natural Language Processing.

## Version

    0.4

## Installation

    go get github.com/christianrondeau/go-wit

## Documentation

* [https://godoc.org/github.com/christianrondeau/go-wit](https://godoc.org/github.com/christianrondeau/go-wit)
adasdas
* [test coverage](http://gocover.io/github.com/christianrondeau/go-wit)
* [lint](http://go-lint.appspot.com/github.com/christianrondeau/go-wit)

## Usage

```go
package main

import (
	"github.com/christianrondeau/go-wit"
	"encoding/json"
	"log"
	"os"
)

func main() {
	client := wit.NewClient(os.Getenv("WIT_ACCESS_TOKEN"))

	// Process a text message
	request := &wit.MessageRequest{}
	request.Query = "Hello world"
	result, err := client.Message(request)
	if err != nil {
		println(err)
		os.Exit(-1)
	}
	log.Println(result)
	data, _ := json.MarshalIndent(result, "", "    ")
	log.Println(string(data[:]))

	// Process an audio/wav message
	request = &wit.MessageRequest{}
	request.File = "../_audio_sample/helloWorld.wav"
	request.ContentType = "audio/wav;rate=8000"
	result, err = client.AudioMessage(request)
	if err != nil {
		println(err)
		os.Exit(-1)
	}
	log.Println(result)
	data, _ = json.MarshalIndent(result, "", "    ")
	log.Println(string(data[:]))
}

// Output:

// structs:
// &{bf699a8f-bc90-4fb4-a715-bd8bd77749db Hello world {hello {{ } []} 0.996}}
// &{54ed4e6d-0653-453e-8c0c-81da57c3846c hello world {hello {{ } []} 0.993}}

// json:
// {
//     "msg_id": "76f1c370-bd92-417f-8cb3-e1419d1a9cb3",
//     "msg_body": "Hello world",
//     "outcome": {
//         "intent": "hello",
//         "entities": {
//             "metric": {},
//             "datetime": null
//         },
//         "confidence": 0.996
//     }
// }
// {
//     "msg_id": "322f9b61-0f75-4953-a392-f8eca058a12f",
//     "msg_body": "hello world",
//     "outcome": {
//         "intent": "hello",
//         "entities": {
//             "metric": {},
//             "datetime": null
//         },
//         "confidence": 0.993
//     }
// }
```

## Testing

Set the `WIT_ACCESS_TOKEN` environment variable to your [Wit](https://wit.ai) Server API token

    cd go-wit
    go test

## License

[MIT](LICENSE.txt)

## Author

Original Author: Jason Goecke [@jsgoecke](http://twitter.com/jsgoecke)
This repository's maintainer: Christian Rondeau [@christianrondeau](http://twitter.com/christianrondeau)
