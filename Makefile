.PHONY: run
run:
	go run .

.PHONY: build
build:
	go build .

.PHONY: test
test:
	go test -v -race -cover ./...

.PHONY: coverage
coverage:
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

.PHONY: bench
bench:
	go test -run=^$$ -bench=. -benchmem ./...

.PHONY: lint
lint:
	golangci-lint run -v --fix ./...

.PHONY: clean
clean:
	go clean
	$(RM) cover.out cover.html

.PHONY: docker
docker:
	docker build -t scaffold:latest .