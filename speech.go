// Copyright (c) 2014 Jason Goecke
// messages.go

package wit

// MessageRequest represents a request to process a message
type MessageRequest struct {
	File         string `json:"file,omitempty"`
	MsgID        string `json:"msg_id,omitempty"`
	Context      string `json:"context, omitempty"`
	ContentType  string `json:"contentType, omitempty"`
	FileContents []byte `json:"-"`
}

// AudioMessage requests processing of an audio message (https://wit.ai/docs/api#toc_8)
//
// 		request := &MessageRequest{}
// 		request.File = "./_audio_sample/helloWorld.wav"
//		request.FileContents = data
//		request.ContentType = "audio/wav;rate=8000"
// 		message, err := client.AudioMessage(request)
func (client *Client) AudioMessage(request *MessageRequest) (*MessageResponse, error) {
	result, err := postFile(client.APIBase+"/speech", request)
	if err != nil {
		return nil, err
	}
	message, err := parseMessage(result)
	if err != nil {
		return nil, err
	}
	return message, nil
}
