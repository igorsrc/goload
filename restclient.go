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

func Post(addr string, contentType string, body *bytes.Buffer) *Result {
	resp, err := http.Post(addr, contentType, body)
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
