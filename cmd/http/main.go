package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/josemyduarte/printer/internal/handler"
)

func main() {
	http.HandleFunc("/", handler.Serve)

	addr := determineListenAddress()
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func determineListenAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("$PORT not set, using :80 as default")
		return ":80"
	}
	return ":" + port
}
