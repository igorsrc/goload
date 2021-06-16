package service

import (
	"bytes"
	"log"
	"net/http"
)

func HttpGet(addr string) *Result {
	resp, err := http.Get(addr)
	if err != nil {
		log.Println(err)
		return &Result{Code: -1, Err: err}
	}
	return &Result{Code: resp.StatusCode}
}

func Post(addr string, contentType string, bodyStr string) *Result {
	body := replaceRegexp(bodyStr)
	resp, err := http.Post(addr, contentType, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Println(err)
		return &Result{Code: -1, Err: err}
	}
	return &Result{Code: resp.StatusCode}
}

func Put(addr string, contentType string, bodyStr string) *Result {
	body := replaceRegexp(bodyStr)
	req, err := http.NewRequest(http.MethodPut, addr, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return &Result{Code: -1, Err: err}
	}

	client := &http.Client{}

	req.Header.Set("Content-type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		return &Result{Code: -1, Err: err}
	}
	return &Result{Code: resp.StatusCode}
}
