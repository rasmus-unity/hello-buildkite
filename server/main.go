package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = 3000

var infoLogger, errorLogger *log.Logger

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	if name == "" {
		name = "Buildkite CI"
	}

	s := fmt.Sprintf("Hi there, %s!", name)
	fmt.Fprintf(w, s)
}

// write error log message, which will trigger Harness.io to rollback deployment
func rollbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Rolling back...")
	log.Println("ROLLBACK!")
	fmt.Fprintf(w, "%s: Ok, wrote rollback log message - harness.io should now rollback to previous version", time.Now().UTC())
}

func main() {
	fmt.Println("Starting Harness.io test app (this is a normal log stdout log message)")

	http.HandleFunc("/", handler)
	http.HandleFunc("/rollback", rollbackHandler)

	fmt.Println(fmt.Sprintf("Server started on port %d", port))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
