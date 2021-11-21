docker-compose build
docker-compose run --rm app make install
docker-compose run --rm app sql-migrate up
docker-compose up -d
docker-compose ps
