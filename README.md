# LiProduce

A marketplace for small farmers to sell fresh produce to high end restaurants.

## Packages

| Package                                                                  | Description                                                                        |
| ------------------------------------------------------------------------ | ---------------------------------------------------------------------------------- |
| [server](https://github.com/SatvikR/liproduce/tree/main/packages/server) | Graphql Server written with apollo, express, typeorm, type-graphql, and postgreSQL |
| [web](https://github.com/SatvikR/liproduce/tree/main/packages/web)       | Website written with Next.js, React, URQL, and typescript                          |

## Run Locally

1. Setup

- In a psql shell

```sql
CREATE DATABASE liproduce;
```

- Environment Variables

```sh
# Create file ./packages/server/.env.dev, change username and password depending on your situation
CORS_ORIGIN=http://localhost:3000
DB=liproduce
DB_USERNAME=postgres
DB_PASSWORD=postgres
PORT=4000
```

1. Install packages

```
yarn
```

1. Start server and website

```
yarn dev
```
