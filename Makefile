start:
	go run server.go

docker:
	docker-compose start

gqlgen:
	go generate ./...