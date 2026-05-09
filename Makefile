.PHONY: vet test lint

vet:
	go vet ./...

test:
	go test ./... -race -count=1

lint:
	@command -v golangci-lint >/dev/null || (echo "golangci-lint missing"; exit 1)
	golangci-lint run
