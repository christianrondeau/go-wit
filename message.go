// Copyright (c) 2014 Jason Goecke
// messages.go

package wit

import (
	"encoding/json"
	"net/url"
)

// MessageResponse represents a Wit message (https://wit.ai/docs/api#toc_3)
type MessageResponse struct {
	MsgID    string                     `json:"msg_id"`
	Text     string                     `json:"_text"`
	Entities map[string][]MessageEntity `json:"entities"`
}

// MessageEntity represents the entity portion of a Wit message
type MessageEntity struct {
	Metadata   *string     `json:"metadata,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	Confidence float64     `json:"confidence"`
	Body       *string     `json:"body,omitempty"`
	Start      *int64      `json:"start,omitempty"`
	End        *int64      `json:"end,omitempty"`
	Type       *string     `json:"type,omitempty"`
	Unit       *string     `json:"unit,omitempty"`
	From       interface{} `json:"from,omitempty"`
	To         interface{} `json:"to,omitempty"`
	Values     interface{} `json:"values,omitempty"`
}

// DatetimeValue represents the datetime value portion of a Wit message
type DatetimeInterval struct {
	From Datetime `json:"from"`
	To   Datetime `json:"to"`
}

type Datetime struct {
	Value string  `json:"value"`
	Grain string  `json:"grain"`
	Type  *string `json:"type"`
}

type AmountOfMoneyInterval struct {
	Unit  *string  `json:"unit"`
	Value *float64 `json:"value"`
}

// Parses the JSON into a Message
//
//		message, err := parseMessage([]byte(data))
func parseMessage(data []byte) (*MessageResponse, error) {
	message := &MessageResponse{}
	err := json.Unmarshal(data, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

// Message requests processing of a text message (https://wit.ai/docs/http/20170307#get--message-link)
//
// 		message, err := client.Message("Hello world")
func (client *Client) Message(q string) (*MessageResponse, error) {
	result, err := get(client.APIBase + "/message?q=" + url.QueryEscape(q))
	if err != nil {
		return nil, err
	}
	message, err := parseMessage(result)
	if err != nil {
		return nil, err
	}
	return message, nil
}
