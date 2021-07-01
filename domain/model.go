package domain

type Request struct {
	Address    string
	Payload    string
	Count      int
	Concurrent int
	Backoff    int
	Debug      bool
}

type Result struct {
	Code int
	Err  error
}
