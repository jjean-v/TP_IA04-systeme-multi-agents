package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// servervame
	addr := ":8080"

	// create file server handler
	fs := http.FileServer(http.Dir("./"))

	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	// start HTTP server with `fs` as the default handler
	log.Fatal(http.ListenAndServe(addr, fs))
}
