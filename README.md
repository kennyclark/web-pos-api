# Web-POS API
This is a RESTful API for a simple web-based POS system.

[![License: MIT (Non-SaaS)](https://img.shields.io/badge/license-MIT--Non--SaaS-blue.svg)](./LICENSE)

## Requirements
1. Download and install [GoLang 1.25+](https://go.dev/doc/install)
2. Download and install [PostgreSQL](https://www.postgresql.org/download/)
3. Install [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation) - for database migrations
4. Install [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html#go-install) - for generating fully type-safe idiomatic Go code from SQL
5. Download and Install [Bruno](https://docs.usebruno.com/get-started/bruno-basics/download) - for API testing

## Environment Variables
The following environment variables are required:
| Variable | Description |
| --- | --- |
| DATABASE_URL | The URL of the PostgreSQL server (*postgres://user:password@host:port/webpos_dev?sslmode=disable*) | 
| CORS_ALLOWED_ORIGINS | The allowed origins |
| MODE | The mode of the server (*development, production*) |
| PORT | The port of the server |

## Installing Dependencies
To install dependencies, run the following command:
```bash
go mod tidy
```

## Database Migrations
To create the PostgreSQL database for development, run the following command:
```bash
psql -h localhost -U postgres -c "CREATE DATABASE webpos_dev;"
```

To run all pending migrations, run the following command:
```bash
migrate -database $DATABASE_URL -path internal/db/migrations up
```

To rollback all migrations, run the following command:
```bash
migrate -database $DATABASE_URL -path internal/db/migrations down
```

To rollback 1 migration, run the following command:
```bash
migrate -database $DATABASE_URL -path internal/db/migrations down 1
```

## Running the Server
To run the server, run the following command:
```bash
go run ./cmd/app/main.go
```
> Note: Run migrations before running the server

## Running the Server with Live Reloading (for development)
To run the server with live reloading, run the following command:
```bash
go tool air
```
> Note: Run migrations before running the server

## License

This project is licensed under the **MIT License (Non-SaaS Variant)**.
You are free to use, modify, and self-host this software for personal or commercial use,
but **you may not offer it as a service or SaaS product** to third parties.

See [LICENSE](./LICENSE) for full details.
