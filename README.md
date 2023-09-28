# url-shortener demo app

This app uses a postgres database running within a docker container. To run the app, you need to run database, copy the schema file into the docker container and execute it, then run the Go app.

## Step 1: Run the database
Run the following commands from the `docker` directory in this repo:
- Build the image
  - `docker build -t url-shortener-postgres .`
- Run it using a user-friendly container name which we need for subsequent steps)
  - `docker run --detach -p 5432:5432 --name="url-shortener-postgres" url-shortener-postgres`
- Copy the db schema file
  - `docker cp ./url.sql url-shortener-postgres:/docker-entrypoint-initdb.d/url.sql`
- Execute it to create the required table
  - `docker exec -u postgres url-shortener-postgres psql urlshortener urlshortener -f docker-entrypoint-initdb.d/url.sql`

## Step 2: Run the app
From the main directory: `source .env && go run cmd/url-shortener/main.go`

## To run tests
go test ./...

