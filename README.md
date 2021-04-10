# liproduce

A platform/marketplace for small farmers to sell fresh produce restaurants

## Tech used

| Name                                                    | Usage                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| [golang](https://golang.org/)                           | Language used for the server                            |
| [postgreSQL](https://www.postgresql.org/)               | Database                                                |
| [react](https://reactjs.org/)                           | Library used on frontend (with typescript)              |
| [nextjs](https://nextjs.org/)                           | React framework used on frontend                        |
| [gqlgen](https://gqlgen.com/)                           | Framework used for the graphql server written in golang |
| [go-pg](https://pg.uptrace.dev/)                        | PostgreSQL ORM used in the server                       |
| [go-pg/migrations](https://github.com/go-pg/migrations) | PostgreSQL migration framework/runner                   |

## To run

1. DB setup

Create your database with

```sql
CREATE DATABASE liproduce_db_name;
```

in a psql shell

Create a file called `.env` in the root of the project with the following variables

```
DB_ADDR=<ex. localhost:5432>
DB_USER=<your db username>
DB_PASSWORD=<your db password>
DB_NAME=<ex. liproducedb>
```

> If you are on windows, run the following steps in git bash or cygwin, and make sure you have make installed

2. Run migrations (required)

```
make migrateinit
make migrateup
# use make migratedown to go down a version
```

3. Start server

```
make start
# To rebuild and start:
make dev
```

Go to http://localhost:8080/ for the graphql playground, http://localhost:8080/query for the graphql endpoint
