# Go Postgres Base

## Pre-reqs
1. Have docker and docker-compose
1. Have psql installed

## One Time Setup (to initialize the DB).
1. `docker-compose up`
2. In another terminal: `psql -U postgres -h localhost -f dbScripts/dbInitialization.sql`
    - Use `postgres` as a password.

