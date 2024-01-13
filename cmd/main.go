package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"goprojs/pulsar"
	"net/http"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		pulsar.ListenAndServe()
	}()
	err := http.ListenAndServe(":8043", nil)
	if err != nil {
		return
	}
}
