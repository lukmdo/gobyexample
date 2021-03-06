package main

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
}
