package main

import (
	"log"

	"github.com/Reimei1213/lab/distributed-services-with-go/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":9000")
	log.Fatal(srv.ListenAndServe())
}
