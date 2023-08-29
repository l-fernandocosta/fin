include .env

c_m:
	# create a new migration
	migrate create -ext sql -dir db/migrations -seq $(name)

dbup:
	# postgres up - creates a postgres server
	docker compose up -d

dbdown:
	# postgres down - delete postgres server
	docker compose down

dbcreate:
	# create a database
	docker compose exec -it postgres createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB_NAME}

dbrmv:
	# remove a database
	docker compose exec -it postgres dropdb --username=${POSTGRES_USER}  ${POSTGRES_DB_NAME}

migrate:
	# run migrations
	migrate -path ${MIGRATION_PATH} -database  ${DB_URL} up
migratedown:
	# down migrations
	migrate -path ${MIGRATION_PATH} -database  ${DB_URL} down