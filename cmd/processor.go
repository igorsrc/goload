package cmd

import (
	"fmt"
	"goload/domain"
	"goload/web"
	"time"
)

type httpFunc func(req *domain.Request) *domain.Result

func Get(req *domain.Request) *domain.Result {
	return web.HttpGet(req.Address, req.Debug)
}

func Post(req *domain.Request) *domain.Result {
	return web.HttpPost(address, token, payload, debug)
}

func Process(req *domain.Request, f httpFunc) string {
	start := time.Now()

	semaphore := make(chan struct{}, req.Concurrent)
	getResult := make(chan *domain.Result, req.Count)
	defer func() {
		close(semaphore)
		close(getResult)
	}()

	for i := 1; i <= req.Count; i++ {
		go func() {
			semaphore <- struct{}{}
			getResult <- f(req) // <-- push result
			time.Sleep(time.Duration(req.Backoff) * time.Millisecond)
			<-semaphore
		}()
	}

	var resArray []domain.Result

	for {
		result := <-getResult // <-- drain result
		resArray = append(resArray, *result)
		if len(resArray) >= req.Count {
			break
		}
	}

	ms := time.Since(start).Milliseconds()
	msg := "%d/%d in %vms."
	return fmt.Sprintf(msg, len(resArray), count, ms)
}
