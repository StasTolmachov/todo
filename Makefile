build:
	docker-compose build todo

run:
	docker-compose up todo

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:QWERTY@0.0.0.0:5436/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go
