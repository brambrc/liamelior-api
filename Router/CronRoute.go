package Router

import (
	"fmt"
	"liamelior-api/Cronjob"

	"github.com/robfig/cron/v3"
)

func CronJob() {
	c := cron.New()

	// Add a cron job with the specified schedule (every day at 18:00 PM)
	_, err := c.AddFunc("0 18 * * *", Cronjob.TriggerAPICronJob)
	if err != nil {
		fmt.Println("Failed to add cron joWb:", err)
		return
	}

	// Start the cron job scheduler
	c.Start()
}
