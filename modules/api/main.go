// Command api is a tiny HTTP service used to exercise the CI pipeline.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler returns the health endpoint handler.
func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprint(w, "ok")
	})
	return mux
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", Handler()))
}
