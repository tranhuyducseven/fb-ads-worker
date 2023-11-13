package job

func RunAllJob() {

	//s := gocron.NewScheduler(time.UTC)
	//
	//s.Cron("*/10 * * * *").Do(func() {
	//
	//	SyncVNPost()
	//})
	//
	//s.StartBlocking()

	SyncVNPost()
}
