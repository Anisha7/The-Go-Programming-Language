// Server1 + also counts the number of requests

// Server2 is a minimal "echo" and counter server

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// We use mu to prevent concurrency bugs with count updates
// If 2 concurrent requests try to update the count at the same time,
// it might be incremented consistently; causing a "race condition" bug
// mu.Lock() and mu.Unlock() make sure that only one goroutine accesses
// the variable at a time
var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
