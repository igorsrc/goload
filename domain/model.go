package domain

type Request struct {
	Address    string
	Payload    string
	Count      int
	Concurrent int
	Backoff    int
}

type Result struct {
	Code   int
	Time   int64
	MadeAt string
}
