migrateUp:
	migrate -path migrations -database "postgres://postgres:qwe@172.18.0.1:5435/goshop?sslmode=disable" up

migrateDown:
	migrate -path migrations -database "postgres://postgres:qwe@172.18.0.1:5435/goshop?sslmode=disable" down