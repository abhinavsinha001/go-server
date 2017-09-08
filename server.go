package main

import (
    "fmt"
    "net/http"
    "github.com/paulbellamy/ratecounter"
    "time"
)

var counter = ratecounter.NewRateCounter(1 * time.Second)
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    counter.Incr(1)
}

func metricsQpsHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Qps %d!", counter.Rate())
}

func main() {
    http.HandleFunc("/application", handler)
    http.HandleFunc("/metrics/qps", metricsQpsHandler)
    http.ListenAndServe(":8080", nil)
}
