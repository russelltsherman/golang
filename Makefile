

deps:
	go install -v golang.org/x/tools/cmd/godoc@latest
.PHONY: deps

doc: deps
	open http://localhost:8001
	godoc -http=:8001
.PHONY: doc

test:
	go test -v -tags all_tests ./...
.PHONY: test

test/algo:
	go test -v -tags algorythm_tests ./...
.PHONY: test/algo

test/data:
	go test -v -tags datastructure_tests ./...
.PHONY: test/data

test/patt:
	go test -v -tags pattern_tests ./...
.PHONY: test/patt

test/bench:
	go test -v -tags all_tests ./... -bench=.
.PHONY: test/bench

test/cover:
	go test -v -tags all_tests ./... -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html
.PHONY: test/cover
