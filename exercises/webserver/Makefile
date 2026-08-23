
stack: stack/up
.PHONY: stack

stack/up:
	docker-compose --file stack.yml up --detach
.PHONY: stack/up

stack/down:
	docker-compose --file stack.yml down
.PHONY: stack/down

doc:
	godoc -http=:8001
	open localhost:8001
.PHONY: doc

test: stack
	go test -v
.PHONY: test

test/coverage: stack
	go test -cover -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html
	open coverage.html
.PHONY: test/coverage

