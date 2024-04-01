generate: 
	go run github.com/99designs/gqlgen

run:
	go run server.go

start-db:
	docker-compose up -d

stop-db:
	docker-compose down
