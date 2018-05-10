package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
	"go.opencensus.io/zpages"
)

var (
	err error

	// metrics
	someKey, _ = tag.NewKey("my.org/keys/someKey")
	//someMetric *stats.Int64Measure
	someMetric = stats.Int64("my.org/measure/testCounter", "Some test counter", stats.UnitDimensionless)
	someTimer  = stats.Int64("my.org/timer/fooTimer", "Some test timer", stats.UnitMilliseconds)
)

func fooFunc(ctx context.Context) {
	fmt.Println("Func")
	ctx, err := tag.New(ctx,
		tag.Insert(someKey, "someKeyValue"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, span := trace.StartSpan(ctx, "my.org/fooFunc")
	defer span.End()

	// Sleep for [1,10] milliseconds to fake work.
	n := time.Duration(rand.Intn(100)+1) * time.Millisecond
	time.Sleep(n)

	stats.Record(ctx, someMetric.M(142))
	stats.Record(ctx, someTimer.M(int64(n)))

	fooFuncHelper(ctx)
}

func fooFuncHelper(ctx context.Context) {
	fmt.Println("Func")
	ctx, span := trace.StartSpan(ctx, "my.org/fooFuncHelper")
	defer span.End()
}

func handleZpages() {
	// http://127.0.0.1:8081/tracez
	// http://127.0.0.1:8081/rpcz
	go func() { log.Fatal(http.ListenAndServe(":8081", zpages.Handler)) }()
}

func main() {
	log.Print("Setup")
	ctx := context.Background()

	handleZpages()

	// register views
	if err := view.Register(&view.View{
		Name:        "my.org/views/fooView",
		Description: "some distribution over time",
		TagKeys:     []tag.Key{someKey},
		Measure:     someTimer,
		Aggregation: view.Distribution(0, 1<<16, 1<<32),
	}); err != nil {
		log.Fatalf("Cannot subscribe to the view: %v", err)
	}

	// setup exporters for subscribed views
	e := &exporter.PrintExporter{}
	view.RegisterExporter(e)
	trace.RegisterExporter(e)

	//PrometheusExporter()
	//StackdriverExporter()

	// always trace for this demo
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	// report stats at every N second.
	view.SetReportingPeriod(3 * time.Second)

	//someKey, err = tag.NewKey("my.org/keys/testKey")
	//if err != nil {
	//	log.Fatal(err)
	//}

	// actual worker
	for i := 0; i < 10; i++ {
		fooFunc(ctx)
	}
	log.Print("Func Done")

	// Wait for a duration longer than reporting duration to ensure the stats
	// library reports the collected data.
	fmt.Println("Wait longer than the reporting duration...")
	time.Sleep(160 * time.Second)

	log.Print("Done")
}
