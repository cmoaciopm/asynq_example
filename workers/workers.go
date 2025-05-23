package main

import (
	"context"
	"log"
	"time"

	"github.com/hibiken/asynq"

	"asynq-quickstart/task"
)

func loggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		start := time.Now()
		log.Printf("Start processing %q", t.Type())
		err := h.ProcessTask(ctx, t)
		if err != nil {
			return err
		}
		log.Printf("Finished processing %q: Elapsed Time = %v", t.Type(), time.Since(start))
		return nil
	})
}

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.Use(loggingMiddleware)
	mux.HandleFunc("email:welcome", task.HandleWelcomeEmailTask)
	mux.HandleFunc("email:reminder", task.HandleReminderEmailTask)

	// Use asynq.HandlerFunc adapter for a handler function
	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
