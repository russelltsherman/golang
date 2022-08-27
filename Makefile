
deps:
	go install -v golang.org/x/tools/cmd/godoc@latest
.PHONY: deps

doc: deps
	open http://localhost:8001
	godoc -http=:8001
.PHONY: doc

test:
	go test -v ./...
.PHONY: test

cover:
	go test -v ./... -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html
.PHONY: test/cover

test/bench:
	go test -v ./... -bench=.
.PHONY: test/bench


