.PHONY: *

test:
	go test ./...

format:
	gofmt -s -w .

crawl:
	go run pkg/crawler/main.go

decrypt:
	gpg --quiet --batch --yes --decrypt --passphrase="abc" --output "dir" "gpg.file"
