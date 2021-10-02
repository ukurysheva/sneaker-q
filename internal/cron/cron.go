package cron

import (
	"github.com/jasonlvhit/gocron"
)

// RunCron is running a cron task every 10 seconds (while on dev mode)
// in production it is supposed to be running every night
func RunCron(task func()) (bool, error) {
	task()
	// q := make(chan bool)
	// go parsejob(q, task)

	// // Close channel because we on dev
	// time.Sleep(18 * time.Second)
	// q <- true
	// close(q)
	// fmt.Println("finish")

	return true, nil
}

func parsejob(quit <-chan bool, task func()) {
	for {
		g := gocron.NewScheduler()
		g.Every(10).Second().Do(task)

		select {
		case <-quit:
			g.Clear()
			return
		case <-g.Start():
		}
	}
}
