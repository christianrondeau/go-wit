// Copyright (c) 2014 Jason Goecke
// intents.go

package wit

import (
	"encoding/json"
)

// Intents represents intents in the Wit API (https://wit.ai/docs/api#toc_13)

type Intent struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Doc         string             `json:"doc"`
	Metadata    string             `json:"metadata"`
	Entities    []IntentEntity     `json:"entities"`
	Expressions []IntentExpression `json:"expressions"`
}

type Intents []Intent

type IntentEntity struct {
	Entity
	Role string `json:"role"`
}

type IntentExpression struct {
	Body               string                   `json:"body"`
	ID                 string                   `json:"id"`
	ExpressionEntities []IntentExpressionEntity `json:"entities"`
}

type IntentExpressionEntity struct {
	Wisp  string `json:"wisp"`
	Value string `json:"value"`
	Start int64  `json:"start"`
	End   int64  `json:"end"`
}

// Intents lists intents configured in the Wit API (https://wit.ai/docs/api#toc_13)
//
//		result, err := client.Intents()
func (client *Client) Intents() (*Intents, error) {
	result, err := get(client.APIBase + "/intents")
	if err != nil {
		return nil, err
	}
	intents, _ := parseIntents(result)
	return intents, nil
}

// Parses the JSON for an Intent
func parseIntents(data []byte) (*Intents, error) {
	intents := &Intents{}
	err := json.Unmarshal(data, intents)
	if err != nil {
		return nil, err
	}
	return intents, nil
}

func (client *Client) IntentShow(intentId string) (*Intent, error) {
	result, err := get(client.APIBase + "/intents/" + intentId)
	if err != nil {
		return nil, err
	}
	intent := &Intent{}

	err = json.Unmarshal(result, intent)
	if err != nil {
		return nil, err
	}

	return intent, nil
}
