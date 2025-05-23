package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"

	"asynq-quickstart/payload"
)

func handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case "email:welcome":
		var p payload.EmailTaskPayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return err
		}
		log.Printf(" [*] Send Welcome Email to User %d", p.UserID)
	case "email:reminder":
		var p payload.EmailTaskPayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return err
		}
		log.Printf(" [*] Send Reminder Email to User %d", p.UserID)
	default:
		return fmt.Errorf("unexpected task type: %s", t.Type())
	}
	return nil
}

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)

	// Use asynq.HandlerFunc adapter for a handler function
	if err := srv.Run(asynq.HandlerFunc(handler)); err != nil {
		log.Fatal(err)
	}
}
