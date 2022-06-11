migrateUp:
	migrate -path migrations -database "postgres://postgres:qwe@localhost:5435/goshop?sslmode=disable" up

migrateDown:
	migrate -path migrations -database "postgres://postgres:qwe@localhost:5435/goshop?sslmode=disable" down