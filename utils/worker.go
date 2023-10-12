package utils

import (
	"os"

	"github.com/hibiken/asynq"
)

var ClientWorker asynq.Client
var JobInspector asynq.Inspector
var RedisClientOpt asynq.RedisClientOpt

func Client() {
	RedisClientOpt = asynq.RedisClientOpt{Addr: os.Getenv("REDIS_HOST"), DB: 1}

	jobInspector := asynq.NewInspector(RedisClientOpt)
	client := asynq.NewClient(RedisClientOpt)
	ClientWorker = *client
	JobInspector = *jobInspector
	// defer client.Close()
}
