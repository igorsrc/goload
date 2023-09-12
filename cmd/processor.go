package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/goload/domain"
	"github.com/goload/web"
)

type httpFunc func(req *domain.Request) *domain.Result

func Get(req *domain.Request) *domain.Result {
	return web.HttpGet(req.Address)
}

func Post(req *domain.Request) *domain.Result {
	return web.HttpPost(address, token, payload)
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
	var report = createFile()
	defer report.Close()
	var resArray []*domain.Result
	for {
		result := <-getResult // <-- drain result
		resArray = append(resArray, result)
		if len(resArray) >= req.Count {
			break
		}
	}

	ms := time.Since(start).Milliseconds()

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Test for url: %s\n", req.Address))
	sb.WriteString(fmt.Sprintf("Time: %d ms.\n", ms))
	sb.WriteString(fmt.Sprintf("Request done: %d\n", len(resArray)))
	sb.WriteString("--------------------------------------------------------------------------------\n")

	var sbReqs strings.Builder
	sbReqs.WriteString("--------------------------------------------------------------------------------\n")
	sbReqs.WriteString("List requests:\n")
	codeCounter := make(map[int]int)
	for _, r := range resArray {
		sbReqs.WriteString(fmt.Sprintf("%s\t code: %d\t time: %d ms.\n", r.MadeAt, r.Code, r.Time))
		codeCounter[r.Code]++
	}
	for k, v := range codeCounter {
		sb.WriteString(fmt.Sprintf("Code %d count: %d\n", k, v))
	}

	_, err := report.WriteString(sb.String() + sbReqs.String())
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("Task done. Report: %s", report.Name())
}

func createFile() *os.File {
	var filename = fmt.Sprintf("out/task_%d.txt", time.Now().UnixMilli())
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return f
}
