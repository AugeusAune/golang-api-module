MIGRATE = migrate
MIGRATION_PATH = database/migration
DB_MIGRATE_URL = $(shell grep DB_MIGRATE_URL .env | cut -d '=' -f2-)

migration_up:
	$(MIGRATE) -path $(MIGRATION_PATH) -database $(DB_MIGRATE_URL) -verbose up

migration_down:
	$(MIGRATE) -path $(MIGRATION_PATH) -database $(DB_MIGRATE_URL) -verbose down $(or $(step), 1)

migration_fix:
	$(MIGRATE) -path $(MIGRATION_PATH) -database $(DB_MIGRATE_URL) force $(version)

migration_create:
	$(MIGRATE) create -ext sql -dir $(MIGRATION_PATH) -seq $(name)

http:
	go run cmd/api/main.go

scheduler:
	go run cmd/scheduler/main.go

queue:
	go run cmd/queue/main.go
