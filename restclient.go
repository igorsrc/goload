package main

import (
	"bytes"
	. "goload/model"
	"log"
	"net/http"
)

func Get(addr string) *Result {
	resp, err := http.Get(addr)
	if err != nil {
		log.Println(err)
		return &Result{
			Code: -1,
			Err:  err,
		}
	}
	return &Result{
		Code:    resp.StatusCode,
		Payload: &resp.Body,
	}
}

func Post(addr string, contentType string, bodyStr string) *Result {
	body := replaceRegexp(bodyStr)
	resp, err := http.Post(addr, contentType, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Println(err)
		return &Result{
			Code: -1,
			Err:  err,
		}
	}
	return &Result{
		Code:    resp.StatusCode,
		Payload: &resp.Body,
	}
}
