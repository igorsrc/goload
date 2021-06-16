package service

type Request struct {
	Address string
	Payload string
}

type Result struct {
	Code int
	Err  error
}
