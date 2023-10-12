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

	// // You can use cron spec string to specify the schedule.
	// entryID, err := scheduler.Register("* * * * *", task)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("registered an entry: %q\n", entryID)

	// You can use "@every <duration>" to specify the interval.

	entryID, err := scheduler.Register("@every 30s", task)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("registered an entry: %q\n", entryID)

	// You can also pass options.
	// entryID, err := scheduler.Register("0 0 * * *", task, asynq.Queue("default"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("registered an entry: %q\n", entryID)

	if err := scheduler.Run(); err != nil {
		log.Fatal(err)
	}

}
