package model

type Request struct {
	Method     string
	Address    string
	Token      string
	Payload    string
	Count      int
	Concurrent uint32
	Backoff    uint32
	Debug      bool
}
