
<p align="center">

  <img src="https://evo.company/wp-content/uploads/2022/04/fntch_logo.svg" width="350" alt="EVO Fintech logo">
</p>

# Project for Trainee EVO Fintech (second stage)
Program accepts CSV file and parses it's data to database. Data can be gotten from database with filters in JSON. CSV example stored in `csv_example/example.csv`.
The author tried to write this project using the principles of Clean Architecture and Dependency Injection.

## Configuration

Configurations stored in `internal/config/config.go`

- database configuration stored in `DBConfig struct`. If you don't use docker compose set `config.DBConfig.Host == "localhost"`. If you use it set `config.DBConfig.Host == "db"`. 
- port number stored in `const PortNumber (default 8000)`.

## Launch

To run the application:

`go run cmd/web/main.go`

When you run the application, it performs migrations to the database.

## Docker & Docker Compose

To build image docker use command:

`docker build -t trans-app .`

To run docker image use command:

`docker run --name=transaction-app -p 8000:8000 trans-app`

To build docker compose use command:

`docker-compose up --build trans-app`

To run docker compose use command:

`docker-compose up trans-app`

## Endpoints

- `/upload-csv` POST method for uploading a CSV file with key `file`.
- `/get-json` GET method to get data from database in JSON format.
- `/get-csv-file` GET method to get data from database in attached CSV file.

With `/get-json` & `/get-csv-file` can be used filters with keys, examples:

| Key               | Example                                        | Note                                   |
|-------------------|------------------------------------------------|----------------------------------------|
| transaction_id    | `transaction_id=1`                             |                                        |
| terminal_id       | `terminal_id=3506,3507`                        | can be more than only one ID           |
| status            | `status=accepted`                              |                                        |
| payment_type      | `payment_type=cash`                            |                                        |
| date_post_from    | `date_post_from=2022-08-12`                    |                                        |
| date_post_to      | `date_post_to=2022-08-23`                      |                                        |
| payment_narrative | `payment_narrative='А11/27123 від 19.11.2020'` | can search by partially specified data |

It can be used none of them, one of them, several filters or all filters at once:

```/get-json?terminal_id=3518,3506,3507&payment_narrative='ослуг А11/27122 від'&date_post_to=2022-08-17&date_post_from=2022-08-13```

## Features of the program
- the program was made with implement Clean Architecture
- Dependency Injection is implemented
- implemented Gratefull shutdown (not yet)
- implemented the use of migrations (when you run the application, it performs migrations to the database)
- To provide the ability to work with large files:
- - parsing is carried out line by line in goroutines
- - implemented adding records to the database by queries for 50 records
- before adding new data to the database, transaction_id is checked for uniqueness.  When a file is received, only records with a unique transaction_id value are entered into the database, and non-unique transaction_id values are returned with an error.

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
