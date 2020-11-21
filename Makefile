# Apply migration
.PHONY: migup
migup:
	migrate -path migration -database "mysql://user:secret@tcp(localhost:3306)/db" -verbose up

# Drop migration
.PHONY: migdown
migdown:
	migrate -path migration -database "mysql://user:secret@tcp(localhost:3306)/db" -verbose down

# Start all services
.PHONY: up
up:
	cd ./dev && docker-compose up -d