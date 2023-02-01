# Go-Short - URL Shortener

A simple URL shortener API built with Go Fiber framework - _Go short, Go fast!_

## Requirements

- [Go v1.19+](https://go.dev/dl/)
- [MySQL v5.7](https://dev.mysql.com/downloads/mysql/5.7.html)
- [Docker v20+](https://docs.docker.com/desktop/)
- [Docker-Compose v1.29+](https://docs.docker.com/compose/install/)

## Setting up the development environment

1. Execute `go mod download` command to download the dependencies
2. Create a .env file in the project root directory with the environment variables according to the _Environment Variables_ section
3. Execute `go run .` command to start the application

## Environment Variables

- **PORT**: Server port number to start listening (e.g. ":8080")
- **MYSQL_DATABASE**: Database name
- **MYSQL_USER**: Database username
- **MYSQL_PASSWORD**: Database user password
- **MYSQL_ROOT_PASSWORD**: Database root password
- **DB_HOST**: Database host name (e.g. local => "localhost:3306", docker => "db")
- **DOMAIN**: Domain name/origin of the application (e.g. "localhost:8080", "goshort.com")

## Build and deploy

### Local environment

1. Execute **build.sh** script to build the executable binary
2. Execute `./dist/go-short-app` command to start the application

### Docker containers

1. Execute `docker-compose up` command to start the containers including the MySQL database container

## How to use the application

## Shorten a link

- Send a POST request to the **/api/link** endpoint with the request body similar to the following:`{"url": "http://google.com"}`

## Using the shortened link

- Send a GET request to the returned shortened link in the response for the above POST request
