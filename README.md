# Learn

## Stack

- GO
- SQLC
- go-migrate

`docker run --name Learn -p 5432:5432 -e POSTGRES_PASSWORD='local-password' -d bitnami/postgresql:17`

## Migrations

New migrations

`migrate create -ext sql -dir ./migrations -seq update_list_table`

Running Migrations

`migrate -database postgres://postgres:local-password@localhost:5432/list?sslmode=disable -path ./migrations up`

Failing Migration

Ex: 'error: Dirty database version x. Fix and force version.'

`migrate -database "postgres://postgres:local-password@localhost:5432/list?sslmode=disable" -path migrations force x-1`
