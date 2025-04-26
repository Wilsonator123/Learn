# Learn

`docker run --name Learn -p 5432:5432 -e POSTGRES_PASSWORD='local-password' -d bitnami/postgresql:17`

`migrate -database postgres://postgres:local-password@localhost:5432/list -path ./cmd/db/migrations up`