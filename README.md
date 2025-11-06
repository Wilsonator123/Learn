# Learn

S - There are too many new tech and skills I want to learn and develop, and its hard to keep track on what I am learning

T - Build a way to manage and learn new skills and technology

A - Build a board like website which allows me to create tasks and move them across the board depending on its status

R - Was able to develop my skills in Go and HTMX whilst having a way to keep track of tasks

## Stack

- GO
- SQLC
- go-migrate

`docker run --name Progress -p 5432:5432 -e POSTGRES_PASSWORD='local-password' -d bitnami/postgresql:17`

## Migrations

New migrations

`migrate create -ext sql -dir ./migrations -seq update_list_table`

Running Migrations

`migrate -database postgres://postgres:local-password@localhost:5432/progress?sslmode=disable -path ./migrations up`

Failing Migration

Ex: 'error: Dirty database version x. Fix and force version.'

`migrate -database "postgres://postgres:local-password@localhost:5432/list?sslmode=disable" -path migrations force x-1`
