package main

// Refs:
// - https://cloud.google.com/go/docs/reference/
// - https://godoc.org/cloud.google.com/go/trace

import (
	"log"

	"contrib.go.opencensus.io/exporter/stackdriver"

	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

const googleProectID = "mytests-1"

func StackdriverExporter() {
	// may need `gcloud auth application-default login` to set correct account for `googleProectID`
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: googleProectID})
	if err != nil {
		log.Fatal(err)
	}

	// Export to Stackdriver Monitoring
	view.RegisterExporter(exporter)
	trace.RegisterExporter(exporter)

	//// Automatically add a Stackdriver trace header to outgoing requests:
	//client := &http.Client{
	//	Transport: &ochttp.Transport{
	//		Propagation: &propagation.HTTPFormat{},
	//	},
	//}
	//_ = client // use client
	//
	//// All outgoing requests from client will include a Stackdriver Trace header.
	//// See the ochttp package for how to handle incoming requests.
}
