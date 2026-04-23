include .env
export

service_run:
	go run ./cmd/api/main.go

#migrate create -ext sql -dir migrations -seq first_create 
table_up:
	migrate -path migrations -database ${CONN_DB} up

table_down:
	migrate -path migrations -database ${CONN_DB} down

table_fix:
	migrate -path migrations -database ${CONN_DB} force 1

run-http-app:
	docker run -d -p 8080:8080 first_image:latest

postgres_up:
	docker run -e POSTGRES_PASSWORD=2906 -p 5432:5432 -v ./out/pgdata:/var/lib/postgresql -d postgres:18-bookworm

