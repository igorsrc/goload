package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "goload/model"
	"net/url"
	"strings"
	"time"
)

const (
	MethodGet     = "GET"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodDelete  = "DELETE"
	MethodPatch   = "PATCH"
	MethodHead    = "HEAD"
	MethodOptions = "OPTIONS"
)

func ProcessRequests(request *Request) []Result {

	// concurrency control channel
	semChan := make(chan struct{}, request.Concurrent)

	// result channel
	resChan := make(chan *Result)

	// close once job is done
	defer func() {
		close(semChan)
		close(resChan)
	}()

	for i := 0; i < request.Count; i++ {

		go func(i int, req *Request) {

			// fill until limit reaches
			semChan <- struct{}{}

			// process request
			if request.Method == MethodGet {
				resChan <- Get(req.Address)
			}
			if request.Method == MethodPost {
				body := replaceRegexp(request.Payload)
				resChan <- Post(req.Address, "application/json", bytes.NewBuffer([]byte(body)))
			}

			// wait before release
			time.Sleep(time.Duration(request.Backoff) * time.Millisecond)

			// push queue
			<-semChan

		}(i, request)
	}

	// result collection
	var resArray []Result

	for {
		// read result and collect
		result := <-resChan
		resArray = append(resArray, *result)

		// stop when all request done
		if len(resArray) >= request.Count {
			break
		}
	}

	return resArray
}

func validateRequest(request *Request) {

	// validate http method
	isMethodCorrect(request.Method)

	// validate address
	isUrlCorrect(request.Address)

	// validate json
	if request.Method == MethodPost {
		isPayloadCorrect(request.Payload)
	}
}

func isMethodCorrect(method string) {
	if len(strings.TrimSpace(method)) < 2 {
		panic("invalid http method")
	}
	methodsAvailable := []string{
		MethodGet,
		MethodPost,
	}
	for _, m := range methodsAvailable {
		if method == m {
			return
		}
	}
	panic("unexpected http method")
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
	if !json.Valid([]byte(payload)) {
		panic("invalid json")
	}
}

func main() {
	req := &Request{
		Method:     "GET",
		Address:    "https://google-translate1.p.rapidapi.com/language/translate/v2/languages",
		Payload:    "json",
		Count:      150,
		Concurrent: 20,
		Backoff:    200,
		Debug:      true,
	}

	validateRequest(req)

	run := func() string {
		start := time.Now()
		results := ProcessRequests(req)
		ms := time.Since(start).Milliseconds()
		msg := "%d/%d in %vms."
		return fmt.Sprintf(msg, len(results), req.Count, ms)
	}

	fmt.Println(run())
}
