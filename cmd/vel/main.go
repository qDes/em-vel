package main

import (
	"em-vel/internal/app/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/halcmd", handlers.RunCmd)
	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal("cannot start server", err)
	}
}
