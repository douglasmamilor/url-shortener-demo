docker build -t url-shortener-postgres .
docker run --detach -p 5432:5432 --name="url-shortener-postgres" url-shortener-postgres
docker cp ./url.sql url-shortener-postgres:/docker-entrypoint-initdb.d/url.sql
docker exec -u postgres url-shortener-postgres psql urlshortener urlshortener -f docker-entrypoint-initdb.d/url.sql
