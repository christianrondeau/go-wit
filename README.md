# go-wit 

[![Build Status](https://api.travis-ci.org/christianrondeau/go-wit.svg)](http://travis-ci.org/christianrondeau/go-wit) [![GoDoc](https://godoc.org/github.com/christianrondeau/go-wit?status.svg)](http://godoc.org/github.com/christianrondeau/go-wit)

A Go library for the [Wit.ai](http://wit.ai) API for Natural Language Processing.

## Version

    0.4

## Installation

    go get -u github.com/christianrondeau/go-wit

## Reference

* [doc](https://godoc.org/github.com/christianrondeau/go-wit)
* [coverage](http://gocover.io/github.com/christianrondeau/go-wit)
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

## Testing

Set the `WIT_ACCESS_TOKEN` environment variable to your [Wit](https://wit.ai) Server API token

    cd go-wit
    go test

## License

[MIT](LICENSE.txt)

## Author

Original Author: Jason Goecke [@jsgoecke](https://twitter.com/jsgoecke)
Maintainer: Christian Rondeau [@christianrondeau](https://github.com/christianrondeau)
