run:
	go run cmd/api/main.go

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

test-health:
	curl http://localhost:8080/health

test-items:
	curl http://localhost:8080/items