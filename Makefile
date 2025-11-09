.PHONY: test coverage examples

test:
	@go test ./...
	
coverage:
	@./scripts/coverage.sh

examples:
	@./examples/run_all.sh
