up:
	docker-compose --env-file $(PWD)/scripts/.env.dev up

psql:
	psql \
		--host localhost \
		--port 5432 \
		--username=postgres \
		--dbname=default

test:
	go test -v ./...

run:
	go run cmd/nalipka/main.go