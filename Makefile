

doc:
	godoc -http=:8001
	open localhost:8001
.PHONY: doc

test:
	go test -v ./...
.PHONY: test

test/benchmark:
	go test -v ./... -bench=.
.PHONY: test/benchmark


test/coverage:
	go test -v ./... -cover -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html
	open coverage.html
.PHONY: test/coverage
