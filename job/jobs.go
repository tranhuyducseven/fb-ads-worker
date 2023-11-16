package job

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

func RunAllJob() {

	s := gocron.NewScheduler(time.UTC)

	s.Cron("5 * * * *").Do(func() {
		fmt.Println("start sync vnpost data")
		SyncVNPost()
	})

	s.Cron("*/3 * * * *").Do(func() {
		SyncVnPostJourney()
	})

	s.StartBlocking()
}
