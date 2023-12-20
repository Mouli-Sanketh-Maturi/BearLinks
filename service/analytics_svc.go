package service

import (
	"bearLinks/datastore"
	"log"
	"sync"
	"time"
)
var analyticsChannel chan<-Analytics

func InitAnalytics() {
	analyticsChannel = startWorkerPool(10)
}

func CaptureAnalytics(shortLink string, ip string, timestamp int64, method string) {
	analytics := Analytics{ShortLink: shortLink, IP: ip, Timestamp: timestamp, Method: method}
	select {
		case analyticsChannel <- analytics:
		case <-time.After(5 * time.Second):
	}
}

func worker(analyticsTask <-chan Analytics, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range analyticsTask {
		_, err := datastore.GetDb().Exec(ctx, "INSERT INTO analytics (shortlink, ipaddress, timestamp, method) VALUES " +
			"($1, $2, $3, $4)", task.ShortLink, task.IP, task.Timestamp, task.Method)
		if err != nil {
			log.Println(err)
		}
	}
}

func startWorkerPool(numWorkers int) chan<- Analytics {
	tasks := make(chan Analytics)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(tasks, &wg)
	}

	return tasks
}