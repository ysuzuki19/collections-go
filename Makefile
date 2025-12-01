.PHONY: test coverage examples

test:
	@go test ./...

generate:
	@go generate ./...
	
coverage:
	@./scripts/coverage.sh

examples:
	@./examples/run_all.sh
