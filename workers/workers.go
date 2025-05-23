package main

import (
	"log"

	"github.com/hibiken/asynq"

	"asynq-quickstart/task"
)

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc("email:welcome", task.HandleWelcomeEmailTask)
	mux.HandleFunc("email:reminder", task.HandleReminderEmailTask)

	// Use asynq.HandlerFunc adapter for a handler function
	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
