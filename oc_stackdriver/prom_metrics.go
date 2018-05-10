package main

// go get github.com/prometheus/client_golang/prometheus

import (
	"go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/stats/view"
	"log"
	"net/http"
)

func PrometheusExporter() {

	exporter, err := prometheus.NewExporter(prometheus.Options{})
	if err != nil {
		log.Fatal(err)
	}
	view.RegisterExporter(exporter)

	log.Println("Serving prometheus /metrics at :9999")
	http.Handle("/metrics", exporter)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
