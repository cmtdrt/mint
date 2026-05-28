.PHONY: fmt fmtcheck vet test ci

fmt:
	gofmt -w .

fmtcheck:
	@files="$$(gofmt -l .)"; \
	if [ -n "$$files" ]; then \
		echo "gofmt required on:"; \
		echo "$$files"; \
		exit 1; \
	fi

vet:
	go vet ./...

test:
	go test -race -v ./...