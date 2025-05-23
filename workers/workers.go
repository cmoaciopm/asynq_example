package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"

	"asynq-quickstart/payload"
)

func sendWelcomeEmail(ctx context.Context, t *asynq.Task) error {
	var p payload.EmailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] Send Welcome Email to User %d", p.UserID)
	return nil
}

func sendReminderEmail(ctx context.Context, t *asynq.Task) error {
	var p payload.EmailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] Send Reminder Email to User %d", p.UserID)
	return nil
}

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc("email:welcome", sendWelcomeEmail)
	mux.HandleFunc("email:reminder", sendReminderEmail)

	// Use asynq.HandlerFunc adapter for a handler function
	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
