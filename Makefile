migrate:
	migrate create -ext sql -dir migrations -seq management

sqlc:
	sqlc generate

.PHONY: migrate sqlc