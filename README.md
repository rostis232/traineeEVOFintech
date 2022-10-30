# Project for Trainee EVO Fintech (second stage)
Program accepts CSV file and parses it's data to database. Data can be gotten from database with filters in JSON. CSV example stored in `csv_example/example.csv`.
The author tried to write this project using the principles of Clean Architecture and Dependency Injection.

## Configuration

Configurations stored in `internal/config/config.go`

- database configuration stored in `DBConfig struct`. Note! Project uses PostgreSQL 15.0.
- port number stored in `const PortNumber (default 8000)`.

## Launch

To run without migration:

`go run cmd/web/main.go`

To run with migration type key -m or --migrate:

`go run cmd/web/main.go -m`

`go run cmd/web/main.go --migrate`

## Endpoints

- `/upload-csv` POST method for uploading a CSV file with key `file`
- `/get-json` GET method to get data from database in JSON. It`s supports filters with keys, examples:
| KEY          | example                                | note                         |
|--------------|----------------------------------------|------------------------------|
|transaction_id|`transaction_id=1`                      |                              |
|terminal_id   |`terminal_id=3506,3507`                 |can be more than only one ID  |
- - `terminal_id=3506`, can be more than only one ID: `terminal_id=3506,3507`
- - `status=accepted`
- - `payment_type=cash`
- - `date_post_from=2022-08-12`
- - `date_post_to=2022-08-23`
- - `payment_narrative='А11/27123 від 19.11.2020'`, can search by partially specified data
- - It can be used none of them, one of them, several filters or all filters at once: `/get-json?terminal_id=3518,3506,3507&payment_narrative='ослуг А11/27122 від'&date_post_to=2022-08-17&date_post_from=2022-08-13`

## Documentation 
Swagger documentation is available on [/swagger/index.html](http://localhost:8000/swagger/index.html)

## Project uses:
- Go version 1.19.2
- PostgreSQL 15.0
- [gin router](https://github.com/gin-gonic/gin)
- [golang migate](https://github.com/golang-migrate/migrate)
- [sqlx](https://github.com/jmoiron/sqlx)
- [pq](https://github.com/lib/pq)
- [gocsv parser](https://github.com/gocarina/gocsv)
- [swaggo swag](https://github.com/swaggo/swag)
