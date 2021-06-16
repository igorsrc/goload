package service

import (
	"encoding/json"
	"net/url"
	"strings"
	"time"
)

func Get(addr string, concur, backoff int, count int) []Result {
	// validation
	isUrlCorrect(addr)

	// channels
	semaphore := make(chan struct{}, concur)
	getResult := make(chan *Result)
	defer func() {
		close(semaphore)
		close(getResult)
	}()

	// goroutines
	for i := 1; i <= count; i++ {
		go func() {
			semaphore <- struct{}{}
			getResult <- HttpGet(addr)
			time.Sleep(time.Duration(backoff) * time.Millisecond)
			<-semaphore
		}()
	}

	var resArray []Result

	// listen
	for {
		result := <-getResult
		resArray = append(resArray, *result)
		if len(resArray) >= count {
			break
		}
	}

	return resArray
}

func isUrlCorrect(_url string) {
	u, err := url.ParseRequestURI(_url)
	if err != nil {
		panic(err)
	}
	if !strings.Contains(u.Scheme, "http") {
		panic("unexpected url scheme")
	}
}

func isPayloadCorrect(payload string) {
	if payload != "" {
		if !json.Valid([]byte(payload)) {
			panic("invalid json")
		}
	}
}
