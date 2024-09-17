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

.PHONY: infra
infra:
	cd infra && docker compose up -d

.PHONY: migrate.new
migrate.new:
	atlas migrate new ${name}

.PHONY: migrate.status
migrate.status:
	atlas migrate status --env local

.PHONY: migrate.apply
migrate.apply:
	atlas migrate apply --env local

.PHONY: migrate.lint
migrate.lint:
	atlas migrate lint --env local --git-base main

.PHONY: migrate.re-hash
migrate.re-hash:
	atlas migrate hash --env local