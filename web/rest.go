package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goload/domain"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HttpGet(addr string, debug bool) *domain.Result {
	start := time.Now()
	resp, err := http.Get(addr)
	var result *domain.Result
	if err != nil {
		log.Println(err)
		result = &domain.Result{Code: -1, Err: err}
	} else {
		result = &domain.Result{Code: resp.StatusCode}
	}
	if debug {
		fmt.Printf("Response: code=%d, time=%d ms.\n", resp.StatusCode, time.Since(start).Milliseconds())
	}
	return result
}

func HttpPost(addr string, token string, bodyStr string, debug bool) *domain.Result {
	start := time.Now()
	var result *domain.Result
	req, err := http.NewRequest("POST", addr, bytes.NewBuffer([]byte(replaceRegexp(bodyStr))))
	if err != nil {
		panic("error creating request: " + err.Error())
	}
	if token != "" {
		req.Header.Add("Authorization", "Bearer"+token)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		result = &domain.Result{Code: -1, Err: err}
	} else {
		result = &domain.Result{Code: resp.StatusCode}
	}
	if debug {
		body, _ := io.ReadAll(resp.Body)
		s := string(body)
		fmt.Printf("Response: code=%d, time=%d ms., body=%s\n", resp.StatusCode, time.Since(start).Milliseconds(), s)
	}
	return result
}

func Put(addr string, contentType string, bodyStr string) *domain.Result {
	body := replaceRegexp(bodyStr)
	var result *domain.Result

	req, err := http.NewRequest(http.MethodPut, addr, bytes.NewBuffer([]byte(body)))
	if err != nil {
		panic("error creating request: " + err.Error())
	}

	client := &http.Client{}
	req.Header.Set("Content-type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		result = &domain.Result{Code: -1, Err: err}
	} else {
		result = &domain.Result{Code: resp.StatusCode}
	}
	return result
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
