#! /bin/sh

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER compose;
	CREATE DATABASE compose;
	GRANT ALL PRIVILEGES ON DATABASE compose TO compose;
EOSQL