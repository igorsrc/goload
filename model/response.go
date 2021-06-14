package model

import "io"

type Result struct {
	Code    int
	Payload *io.ReadCloser
	Err     error
}
