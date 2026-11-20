lint:
	golangci-lint run ./... --fix -c .golangci.yaml

fmt:
	go tool swag fmt ./...
	go tool gofumpt -w .

up:
	docker-compose up -d --build

down:
	docker-compose down

create-migration:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make create-migration name=<migration-name>"; \
		exit 1; \
	fi
	goose -dir ./migrations create "$(name)" sql

generate-swagger:
	go tool swag init -g cmd/app/main.go -o api/docs/ --parseDependency --parseInternal --useStructName