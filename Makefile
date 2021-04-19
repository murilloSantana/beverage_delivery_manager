build:
	go build -v -o bin/beverage_delivery_manager ./cmd/main.go

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
	go run -race ./cmd/main.go

generate:
	@go generate ./...

run-docker:
	cd docker && docker-compose up -d --build

stop-docker:
	cd docker && docker-compose stop

remove-docker:
	cd docker && docker-compose down --remove-orphans

gqlgen: delete-generated-resolver
	@go run github.com/99designs/gqlgen --verbose
	@echo "============= Resolve changes ============="
	@git diff handler/graph/generated/resolver.go

delete-generated-resolver:
	@rm -rf ./handler/graph/generated/resolver.go