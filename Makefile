.PHONY: fmt fmtcheck vet test ci

fmt:
	gofmt -w .

fmtcheck:
	test -z "$$(gofmt -l .)"

vet:
	go vet ./...

test:
	go test -race -v ./...

ci: fmtcheck vet test