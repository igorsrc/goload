package main

import (
	"os"

	"github.com/goload/cmd"
)

func main() {
	os.Mkdir("out/", os.ModePerm)
	cmd.Execute()

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
