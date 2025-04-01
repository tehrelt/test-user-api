ifneq (,$(wildcard ./.env))
    include .env
    export
endif

build: gen 
	go build -o ./bin/app.exe ./cmd/app

gen: wire

clean:
	rm -rf ./internal/app/wire_gen.go

wire: 
	wire ./internal/app

migrate.up:
	migrate -path ./migrations -database 'postgres://$(PG_USER):$(PG_PASS)@$(PG_HOST):$(PG_PORT)/$(PG_NAME)?sslmode=disable' up

migrate.down:
	migrate -path ./migrations -database 'postgres://$(PG_USER):$(PG_PASS)@$(PG_HOST):$(PG_PORT)/$(PG_NAME)?sslmode=disable' down

