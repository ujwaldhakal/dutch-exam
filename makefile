.PHONY: *

test:
	go test ./...

format:
	gofmt -s -w .