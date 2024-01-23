install:
	@make build
	@make up
build:
	docker compose build
up:
	docker compose up -d
stop:
	docker compose stop
down:
	docker compose down --remove-orphans
down-v:
	docker compose down --remove-orphans --volumes
restart:
	@make down
	@make up
go:
	docker compose exec go bash
main:
	docker compose exec go bash -c 'go run /go/src/app/main.go'
