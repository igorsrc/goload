package main

import (
	"goload/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	//req := &Request{
	//	Method:     "GET",
	//	Address:    "https://google-translate1.p.rapidapi.com/language/translate/v2/languages",
	//	Payload:    "json",
	//	Count:      150,
	//	Concurrent: 20,
	//	Backoff:    200,
	//	Debug:      true,
	//}
}
