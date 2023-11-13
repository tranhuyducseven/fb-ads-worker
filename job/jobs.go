package job

import (
	"github.com/go-co-op/gocron"
	"time"
)

func RunAllJob() {

	s := gocron.NewScheduler(time.UTC)

	s.Cron("0 * * * *").Do(func() {
		SyncVNPost()
	})

	s.Cron("* * * * *").Do(func() {
		SyncVnPostJourney()
	})

	s.StartBlocking()
}
