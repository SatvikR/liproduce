import { ApolloServer } from "apollo-server-express";
import cors from "cors";
import express from "express";
import "reflect-metadata";
import { buildSchema } from "type-graphql";
import { createConnection } from "typeorm";
import { CORS_ORIGIN, DB_CONFIG, PORT } from "./constants";
import { ProducerResolver } from "./resolvers/producer";

(async () => {
  await createConnection(DB_CONFIG);

  const app = express();

  app.use(
    cors({
      origin: CORS_ORIGIN,
    })
  );

  const apolloServer = new ApolloServer({
    schema: await buildSchema({
      resolvers: [ProducerResolver],
      validate: false,
    }),
  });

  apolloServer.applyMiddleware({
    app,
    cors: false,
  });

  app.listen(PORT, () => {
    console.log(`Server listening on port ${PORT}`);
  });
})().catch((err) => console.log(err));
