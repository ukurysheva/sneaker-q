package cron

func RunCron(task func()) (bool, error) {
	// gocron.Every(6).Hours().Do(task)
	// <-gocron.Start()
	task()

	return true, nil
}
