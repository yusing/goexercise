.PHONY: check test clean

check: test

test:
	@echo "Running all tests..."
	@go test -v ./...

clean:
	@go clean -testcache
