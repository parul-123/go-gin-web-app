# Go Web Application
This project helps in handling Article data integrated with postgres DB, redis with user authentication feature written in Go.

## Pre-requisites
1. Install postgreSQL and create db name - "articles_go_db"
    ```bash
    $ psql -U postgres
    postgres=> create database articles_go_db;
    ```
2. Install redis and start redis-server
   ```bash
   $ redis-server --daemonize yes
   ```
## Run web application
1. Build app
   ```bash
   $ go build -o app
   ```
2. Run app
   ```bash
   $ ./app
   ```
