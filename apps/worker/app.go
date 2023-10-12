package main

import (
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/tasks/handler"
	"api_tinggal_nikah/tasks/tasktype"
	"api_tinggal_nikah/utils"
	"log"

	"github.com/hibiken/asynq"
)

func init() {
	utils.Client()
	db.InitRedisConnection()
	db.InitDB()
}

func main() {

	redisConnection := utils.RedisClientOpt

	// Create and configuring Asynq worker server.
	worker := asynq.NewServer(redisConnection, asynq.Config{
		// Specify how many concurrent workers to use.
		Concurrency: 10,
		// Specify multiple queues with different priority.
		Queues: map[string]int{
			"critical": 6, // processed 60% of the time
			"default":  3, // processed 30% of the time
			"low":      1, // processed 10% of the time
		},
	})

	// Create a new task's mux instance.
	mux := asynq.NewServeMux()

	// Define a task handler for the welcome email task.
	mux.HandleFunc(
		tasktype.DeletedImage,          // task type
		handler.HandlerDeleteImageTask, // handler function
	)

	// Run worker server.
	if err := worker.Run(mux); err != nil {
		log.Fatal(err)
	}

}
