include .env

.PHONY: postgres adminer migrate run

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=${DATABASE_PASSWORD} ${DATABASE_NAME}

adminer:
	docker run --rm -ti --network host adminer

migrate:
	migrate -source file://migrations \
	  -database ${DATABASE_URL} up

run:
	go run cmd/books-go/main.go
