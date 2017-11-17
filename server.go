package main

import (
    "fmt"
    "flag"
    "log"
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
var httpCounter = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "hits_total",
        Help: "Number of http requests.",
    })

func handler(w http.ResponseWriter, r *http.Request) {
   
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    httpCounter.Inc()
}


func main() {    
    flag.Parse()
    
    err := prometheus.Register(httpCounter)
    if err != nil {
    fmt.Println("Push counter couldn't be registered AGAIN, no counting will happen:", err)
        return
    }

    http.Handle("/metrics", promhttp.Handler())
    http.HandleFunc("/", handler)
    //log.Fatal(http.ListenAndServe(*addr, nil))

    log.Fatal(http.ListenAndServe(*addr, nil))
}
