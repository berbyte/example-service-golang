package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "Hello from host: %q The current  time is %q", hostname, time.Now().String())
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
