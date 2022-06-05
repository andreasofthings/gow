package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
   httpPort := os.Getenv("HTTP_PORT")
   if httpPort == "" {
		httpPort = "8080"
	}
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}
