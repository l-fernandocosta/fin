c_m:
	# create a new migration
	migrate create -ext sql -dir db/migrations -seq $(name)
dbup:
	# postgres up - creates a postgres server
	docker compose up -d
dbdown:
	# postgres down - delete postgres server
	docker compose down