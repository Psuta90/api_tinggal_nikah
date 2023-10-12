package tasks

import (
	"api_tinggal_nikah/tasks/tasktype"
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
)

func NewTasksDeletedImage() (*asynq.Task, error) {

	payload, err := json.Marshal(tasktype.DeletedImagePayload{Time: time.Now()})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(tasktype.DeletedImage, payload, asynq.MaxRetry(0)), nil

}
