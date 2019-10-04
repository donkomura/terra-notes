package functions

import (
	"context"
	"encoding/json"
	"log"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

type Logs struct {
	PlanResult  string `json:"planResult"`
	ApplyResult string `json:"applyResult"`
}

func LogToSlack(cxt context.Context, m PubSubMessage) error {
	var logs Logs
	err := json.Unmarshal(m.Data, &logs)
	if err != nil {
		log.Fatalf("Error: %T message: %v", err, logs)
	}
	return nil
}
