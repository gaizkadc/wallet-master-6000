lint:
	golangci-lint run

test:
	go test -race `go list ./...`

run-local-migrations:
	migrate -source file://internal/storage/migrations -database postgres://postgres:secur3passw0rd@localhost:5432/WM6000?sslmode=disable up

docker-up:
	docker-compose -f docker/docker-compose.yaml -p wm6000 up -d --build

docker-down:
	docker-compose -f docker/docker-compose.yaml -p wm6000 down

docker-ps:
	docker-compose -f docker/docker-compose.yaml -p wm6000 ps

docker-logs:
	docker logs -f wm6000