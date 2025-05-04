FROM --platform=linux/amd64 golang:tip-alpine

WORKDIR /app

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /dist/api ./cmd/api

EXPOSE 1323

CMD ["/dist/api"]