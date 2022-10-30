# Project for Trainee EVO Fintech (second stage)
Program accepts CSV file and parses it`s data to data base. Data can be gotten from data base with filters in JSON.

## Configuration

Configurations stored in `internal/config/config.go`

First Launch needs to do some configuration:

- Data Base configuration stored in `DBConfig struct`. Note! Project uses PostgreSQL 15.0 
- port number stored in `const PortNumber (default 8000)`

## Launch

To run without migration:

`go run cmd/web/main.go`

To run with migration type key -m or --migrate:

`go run cmd/web/main.go -m`

`go run cmd/web/main.go --migrate`

## Endpoints

- `/upload-csv` POST method for uploading a CSV file
- `/get-json` GET method to get data from database in JSON. It`s supports filters with fields, examples:
- - `transaction_id=1`
- - `terminal_id=3506`, can be more than only one ID: `terminal_id=3506,3507`
- - `status=accepted`
- - `payment_type=cash`
- - `date_post_from=2022-08-12`
- - `date_post_to=2022-08-23`
- - `payment_narrative='А11/27123 від 19.11.2020'`, can search by partially specified data

## Project uses:
- Go version 1.19.2
- PostgreSQL 15.0
- [gin router](https://github.com/gin-gonic/gin)
- [golang migate](https://github.com/golang-migrate/migrate)
- [sqlx](https://github.com/jmoiron/sqlx)
- [pq](https://github.com/lib/pq)
- [gocsv parser](https://github.com/gocarina/gocsv)
