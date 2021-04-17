build:
	go build -v -o bin/beverage_delivery_manager

test:
	go test -v -race  --cover ./...

bench:
	go test -bench=. ./...

test-coverage:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

format-files:
	go fmt ./...

lint:
	golangci-lint run

run: format-files lint
	go run -race main.go

generate:
	@go generate ./...

run-docker:
	docker-compose up --build

gqlgen: delete-generated-resolver
	@go run github.com/99designs/gqlgen --verbose
	@echo "============= Resolve changes ============="
	@git diff graph/generated/resolver.go

delete-generated-resolver:
	@rm -rf ./graph/generated/resolver.go