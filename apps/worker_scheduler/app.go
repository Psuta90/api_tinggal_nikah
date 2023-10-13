package main

import (
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/tasks"
	"api_tinggal_nikah/utils"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func init() {
	utils.Client()
	db.InitRedisConnection()
}

func main() {

	redisConnection := utils.RedisClientOpt

	// Example of using America/Los_Angeles timezone instead of the default UTC timezone.
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	scheduler := asynq.NewScheduler(
		redisConnection,
		&asynq.SchedulerOpts{
			Location: loc,
		},
	)

	task, err := tasks.NewTasksDeletedImage()
	if err != nil {
		log.Fatalln(err)
	}

	entryID, err := scheduler.Register("0 0 * * *", task)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("registered an entry: %q\n", entryID)

	if err := scheduler.Run(); err != nil {
		log.Fatal(err)
	}

}
