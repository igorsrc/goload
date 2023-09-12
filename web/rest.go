package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/goload/domain"
)

func HttpGet(addr string) *domain.Result {
	start := time.Now()
	resp, err := http.Get(addr)
	if err != nil {
		panic(err)
	}
	return &domain.Result{Code: resp.StatusCode,
		Time: time.Since(start).Milliseconds(), MadeAt: start.Format(time.RFC3339Nano)}
}

func HttpPost(addr string, token string, bodyStr string) *domain.Result {
	start := time.Now()
	req, err := http.NewRequest("POST", addr, bytes.NewBuffer([]byte(replaceRegexp(bodyStr))))
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return &domain.Result{Code: resp.StatusCode,
		Time: time.Since(start).Milliseconds(), MadeAt: start.Format(time.RFC3339Nano)}
}

func Put(addr string, contentType string, bodyStr string) *domain.Result {
	start := time.Now()
	req, err := http.NewRequest(http.MethodPut, addr, bytes.NewBuffer([]byte(replaceRegexp(bodyStr))))
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	req.Header.Set("Content-type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return &domain.Result{Code: resp.StatusCode,
		Time: time.Since(start).Milliseconds(), MadeAt: start.Format(time.RFC3339Nano)}
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
