PROJECT := "mytests-1"

run:
	@go run *.go

open_prom_metrics:
	open http://127.0.0.1:9999/metrics

open_stackdriver:
	open https://app.google.stackdriver.com/settings/accounts/?project=$(PROJECT)

open_stackdriver_traces:
	open https://console.cloud.google.com/traces/traces?project=$(PROJECT)

open_zpages:
	open http://127.0.0.1:8081/tracez

check:
	go fmt $$(go list ./... | grep -v /vendor/)
	go vet $$(go list ./... | grep -v /vendor/)

gometalinter: $(METALINTER_GO)
	time gometalinter --tests --vendor --deadline=60s ./...


