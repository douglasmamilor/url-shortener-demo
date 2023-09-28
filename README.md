# url-shortener demo app

This app uses a postgres database running within a docker container. To run the app, you need to run database, copy the schema file into the docker container and execute it, then run the Go app.

## Step 1: Run the database
Run the following commands from the `docker` directory in this repo:
`docker build -t url-shortener-postgres .` (build the image)
`docker run --detach -p 5432:5432 --name="url-shortener-postgres" url-shortener-postgres` (run it using a user-friendly container name which we need for subsequent steps)
`docker cp ./url.sql url-shortener-postgres:/docker-entrypoint-initdb.d/url.sql` (copy the db schema file)
`docker exec -u postgres url-shortener-postgres psql urlshortener urlshortener -f docker-entrypoint-initdb.d/url.sql` (execute it to create the required table)

## Step 2: Run the app
From the main directory: `go run cmd/url-shortener/main.go`
